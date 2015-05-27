package ly.stealth.mesos.logging

import ly.stealth.mesos.logging.Util.Str
import org.apache.log4j.Logger
import org.apache.mesos.Protos._
import org.apache.mesos.{Executor, ExecutorDriver, Protos}

trait ExecutorBase extends Executor {
  private var shutdownFlag = false
  private val lock = new Object()

  private val logger = Logger.getLogger(this.getClass)

  def parseExecutorConfig(args: Array[String]): ExecutorConfigBase = {
    val parser = new scopt.OptionParser[ExecutorConfigBase]("executor") {
      override def errorOnUnknownArgument = false

      opt[Int]('p', "port").optional().text("Port to bind to.").action { (value, config) =>
        config.copy(port = value)
      }

      opt[String]('c', "producer.config").required().text("Producer config file name.").action { (value, config) =>
        config.copy(producerConfig = value)
      }

      opt[String]('t', "topic").required().text("Topic to produce transformed data to.").action { (value, config) =>
        config.copy(topic = value)
      }
    }

    parser.parse(args, ExecutorConfigBase()) match {
      case Some(config) => config
      case None => sys.exit(1)
    }
  }

  protected def start()

  override def registered(driver: ExecutorDriver, executorInfo: ExecutorInfo, framework: FrameworkInfo, slave: SlaveInfo) {
    logger.info("[registered] framework:" + Str.framework(framework) + " slave:" + Str.slave(slave))
  }

  override def shutdown(driver: ExecutorDriver) {
    logger.info("[shutdown]")

    lock.synchronized {
      shutdownFlag = true
      lock.notifyAll()
    }
  }

  override def disconnected(driver: ExecutorDriver) {
    logger.info("[disconnected]")
  }

  override def killTask(driver: ExecutorDriver, id: TaskID) {
    logger.info("[killTask] " + id.getValue)

    lock.synchronized {
      shutdownFlag = true
      lock.notifyAll()
    }
  }

  override def reregistered(driver: ExecutorDriver, slave: SlaveInfo) {
    logger.info("[reregistered] " + Str.slave(slave))
  }

  override def error(driver: ExecutorDriver, message: String) {
    logger.info("[error] " + message)
  }

  override def frameworkMessage(driver: ExecutorDriver, data: Array[Byte]) {
    logger.info("[frameworkMessage] " + new String(data))
  }

  override def launchTask(driver: ExecutorDriver, task: TaskInfo) {
    logger.info("[launchTask] " + Str.task(task))

    val runStatus = TaskStatus.newBuilder().setTaskId(task.getTaskId).setState(Protos.TaskState.TASK_RUNNING).build

    driver.sendStatusUpdate(runStatus)

    new Thread {
      override def run() {
        var failed = false
        try {
          logger.info("[Started task] " + Str.task(task))
          ExecutorBase.this.start()

          while (!shutdownFlag) {
            lock.synchronized(lock.wait())
          }
        } catch {
          case t: Throwable =>
            logger.warn("", t)
            failed = true
        } finally {
          val finishedStatus = TaskStatus.newBuilder().setTaskId(task.getTaskId)
            .setState(if (failed) Protos.TaskState.TASK_FAILED else Protos.TaskState.TASK_FINISHED).build
          driver.sendStatusUpdate(finishedStatus)
          driver.stop()
        }
      }
    }.start()
  }
}

case class ExecutorConfigBase(port: Int = 0, producerConfig: String = "", topic: String = "")