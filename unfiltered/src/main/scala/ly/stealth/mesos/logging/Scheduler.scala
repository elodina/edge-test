package ly.stealth.mesos.logging

import java.util.UUID

import org.apache.mesos.Protos
import org.apache.mesos.Protos._

object Scheduler extends SchedulerBase {
  private var schedulerConfig: SchedulerConfigBase = null

  override protected def config: SchedulerConfigBase = schedulerConfig

  def main(args: Array[String]) {
    schedulerConfig = parseSchedulerConfig(args)

    start(schedulerConfig, "Unfiltered LogLine Transform Framework")
  }

  override def launchTask(offer: Offer): Option[TaskInfo] = {
    val cpus = getScalarResources(offer, "cpus")
    val mems = getScalarResources(offer, "mem")
    val ports = getRangeResources(offer, "ports")
    val portOpt = ports.headOption.map(_.getBegin)

    if (cpus > schedulerConfig.cpuPerTask && mems > schedulerConfig.memPerTask && portOpt.nonEmpty) {
      val id = "transform-" + UUID.randomUUID().toString
      val taskId = TaskID.newBuilder().setValue(id).build()
      val taskInfo = TaskInfo.newBuilder().setName(taskId.getValue).setTaskId(taskId).setSlaveId(offer.getSlaveId)
        .setExecutor(this.createExecutor(id, portOpt.get))
        .addResources(Protos.Resource.newBuilder().setName("cpus").setType(Protos.Value.Type.SCALAR).setScalar(Protos.Value.Scalar.newBuilder().setValue(schedulerConfig.cpuPerTask)))
        .addResources(Protos.Resource.newBuilder().setName("mem").setType(Protos.Value.Type.SCALAR).setScalar(Protos.Value.Scalar.newBuilder().setValue(schedulerConfig.memPerTask)))
        .addResources(Protos.Resource.newBuilder().setName("ports").setType(Protos.Value.Type.RANGES).setRanges(
        Protos.Value.Ranges.newBuilder().addRange(Protos.Value.Range.newBuilder().setBegin(portOpt.get).setEnd(portOpt.get))
      )).build

      Some(taskInfo)
    } else None
  }

  private def createExecutor(id: String, port: Long): ExecutorInfo = {
    val path = this.schedulerConfig.executor.split("/").last
    val producerConfigPath = this.schedulerConfig.producerConfig.split("/").last
    val cmd = s"java -cp ${this.schedulerConfig.executor} ly.stealth.mesos.logging.Executor " +
      s"--producer.config ${this.schedulerConfig.producerConfig} --topic ${this.schedulerConfig.topic} --port $port"
    ExecutorInfo.newBuilder().setExecutorId(ExecutorID.newBuilder().setValue(id))
      .setCommand(CommandInfo.newBuilder()
      .addUris(CommandInfo.URI.newBuilder.setValue(s"http://${this.schedulerConfig.artifactServerHost}:${this.schedulerConfig.artifactServerPort}/resource/$path"))
      .addUris(CommandInfo.URI.newBuilder.setValue(s"http://${this.schedulerConfig.artifactServerHost}:${this.schedulerConfig.artifactServerPort}/resource/$producerConfigPath"))
      .setValue(cmd))
      .setName("LogLine Transform Executor")
      .setSource("cisco")
      .build
  }
}