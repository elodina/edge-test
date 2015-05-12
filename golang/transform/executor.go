/* Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

package transform

import (
    "github.com/mesos/mesos-go/executor"
    mesos "github.com/mesos/mesos-go/mesosproto"
    "fmt"
    "time"
)

type TransformExecutor struct {
    
}

// Creates a new TransformExecutor with a given config.
func NewTransformExecutor() *TransformExecutor {
    return &TransformExecutor{
    }
}

// mesos.Executor interface method.
// Invoked once the executor driver has been able to successfully connect with Mesos.
// Not used by TransformExecutor yet.
func (this *TransformExecutor) Registered(driver executor.ExecutorDriver, execInfo *mesos.ExecutorInfo, fwinfo *mesos.FrameworkInfo, slaveInfo *mesos.SlaveInfo) {
    fmt.Printf("Registered Executor on slave %s\n", slaveInfo.GetHostname())
}

// mesos.Executor interface method.
// Invoked when the executor re-registers with a restarted slave.
func (this *TransformExecutor) Reregistered(driver executor.ExecutorDriver, slaveInfo *mesos.SlaveInfo) {
    fmt.Printf("Re-registered Executor on slave %s\n", slaveInfo.GetHostname())
}

// mesos.Executor interface method.
// Invoked when the executor becomes "disconnected" from the slave.
func (this *TransformExecutor) Disconnected(executor.ExecutorDriver) {
    fmt.Println("Executor disconnected.")
}

// mesos.Executor interface method.
// Invoked when a task has been launched on this executor.
func (this *TransformExecutor) LaunchTask(driver executor.ExecutorDriver, taskInfo *mesos.TaskInfo) {
    fmt.Printf("Launching task %s with command %s\n", taskInfo.GetName(), taskInfo.Command.GetValue())

    runStatus := &mesos.TaskStatus{
        TaskId: taskInfo.GetTaskId(),
        State:  mesos.TaskState_TASK_RUNNING.Enum(),
    }

    if _, err := driver.SendStatusUpdate(runStatus); err != nil {
        fmt.Printf("Failed to send status update: %s\n", runStatus)
    }

    go func() {
        time.Sleep(30 * time.Second)

        // finish task
        fmt.Printf("Finishing task %s\n", taskInfo.GetName())
        finStatus := &mesos.TaskStatus{
            TaskId: taskInfo.GetTaskId(),
            State:  mesos.TaskState_TASK_FINISHED.Enum(),
        }
        if _, err := driver.SendStatusUpdate(finStatus); err != nil {
            fmt.Printf("Failed to send status update: %s\n", finStatus)
        }
        fmt.Printf("Task %s has finished\n", taskInfo.GetName())
    }()
}

// mesos.Executor interface method.
// Invoked when a task running within this executor has been killed.
func (this *TransformExecutor) KillTask(_ executor.ExecutorDriver, taskId *mesos.TaskID) {
    fmt.Println("Kill task")
}

// mesos.Executor interface method.
// Invoked when a framework message has arrived for this executor.
func (this *TransformExecutor) FrameworkMessage(driver executor.ExecutorDriver, msg string) {
    fmt.Printf("Got framework message: %s\n", msg)
}

// mesos.Executor interface method.
// Invoked when the executor should terminate all of its currently running tasks.
func (this *TransformExecutor) Shutdown(executor.ExecutorDriver) {
    fmt.Println("Shutting down the executor")
}

// mesos.Executor interface method.
// Invoked when a fatal error has occured with the executor and/or executor driver.
func (this *TransformExecutor) Error(driver executor.ExecutorDriver, err string) {
    fmt.Printf("Got error message: %s\n", err)
}
