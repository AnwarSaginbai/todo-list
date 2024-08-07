<script>
  import { PostTask, GetAllTasks, UpdateTask, DeleteTask } from "../wailsjs/go/main/App.js";

  let tasks = [];
  let newTask = "";
  let showTasks = false;
  let successMessage = "";
  let errorMessage = "";
  let taskToDelete = null;

  function addTask() {
    if (newTask.trim()) {
      PostTask({ task: newTask, done: false })
        .then(() => {
          newTask = "";
          successMessage = "Task successfully added!";
          errorMessage = "";
          getAllTasks(); // Refresh tasks list after adding a task
        })
        .catch((error) => {
          errorMessage = "Failed to add task.";
          console.error("Error adding task:", error);
        });
    } else {
      errorMessage = "Task should not be empty.";
    }
  }

  function markTaskAsDone(taskId) {
    UpdateTask(taskId, { task: tasks.find(task => task.id === taskId).task, done: true })
      .then(() => {
        successMessage = "Task marked as done!";
        errorMessage = "";
        getAllTasks(); // Refresh tasks list after updating a task
      })
      .catch((error) => {
        errorMessage = "Failed to update task status.";
        console.error("Error updating task status:", error);
      });
  }

  function deleteTask() {
    if (taskToDelete !== null) {
      DeleteTask(taskToDelete)
        .then(() => {
          successMessage = "Task successfully deleted!";
          errorMessage = "";
          taskToDelete = null; // Reset taskToDelete
          getAllTasks(); // Refresh tasks list after deleting a task
        })
        .catch((error) => {
          errorMessage = "Failed to delete task.";
          console.error("Error deleting task:", error);
        });
    }
    closeModal();
  }

  function openModal(taskId) {
    taskToDelete = taskId;
    document.getElementById('confirm-modal').style.display = 'block';
  }

  function closeModal() {
    document.getElementById('confirm-modal').style.display = 'none';
    taskToDelete = null; // Reset taskToDelete when closing the modal
  }

  function viewTasks() {
    if (showTasks) {
      showTasks = false;
    } else {
      GetAllTasks()
        .then((result) => {
          tasks = result;
          showTasks = true;
          successMessage = "";
          errorMessage = "";
        })
        .catch((error) => {
          errorMessage = "Failed to fetch tasks.";
          console.error("Error fetching tasks:", error);
        });
    }
  }

  function getAllTasks() {
    GetAllTasks()
      .then((result) => {
        tasks = result;
      })
      .catch((error) => {
        errorMessage = "Failed to fetch tasks.";
        console.error("Error fetching tasks:", error);
      });
  }
</script>

<h3>Task Management API</h3>
<div>
  <input type="text" bind:value={newTask} placeholder="Enter a new task" />
  <button class="btn" on:click={addTask}>
    Add Task
  </button>
  <button class="btn" on:click={viewTasks}>
    {showTasks ? 'Hide Tasks' : 'View Tasks'}
  </button>
</div>

{#if showTasks}
  <ul>
    {#each tasks as task}
      <li class={task.done ? 'done-task' : ''}>
        {task.task} - {task.done ? "Done" : "Pending"}
        {#if !task.done}
          <button class="btn" on:click={() => markTaskAsDone(task.id)}>
            Mark as Done
          </button>
        {/if}
        <button class="btn" on:click={() => openModal(task.id)}>
          Delete
        </button>
      </li>
    {/each}
  </ul>
{/if}

{#if successMessage}
  <p class="success-message">{successMessage}</p>
{/if}
{#if errorMessage}
  <p class="error-message">{errorMessage}</p>
{/if}

<!-- Modal -->
<div id="confirm-modal" class="modal">
  <div class="modal-content">
    <span class="close" on:click={closeModal}>&times;</span>
    <p>Are you sure you want to delete this task?</p>
    <button class="btn" on:click={deleteTask}>Yes, Delete</button>
    <button class="btn" on:click={closeModal}>Cancel</button>
  </div>
</div>

<style>
  .btn:focus {
    border-width: 3px;
  }

  input {
    padding: 8px;
    font-size: 16px;
    margin-right: 8px;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    padding: 8px 0;
    border-bottom: 1px solid #ddd;
  }

  .done-task {
    color: green;
    text-decoration: line-through;
  }

  .success-message {
    color: green;
    font-weight: bold;
  }

  .error-message {
    color: red;
    font-weight: bold;
  }

  .modal {
    display: none; /* Hidden by default */
    position: fixed; /* Stay in place */
    z-index: 1; /* Sit on top */
    left: 0;
    top: 0;
    width: 100%; /* Full width */
    height: 100%; /* Full height */
    overflow: auto; /* Enable scroll if needed */
    background-color: rgb(0,0,0); /* Fallback color */
    background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
  }

  .modal-content {
    background-color: #fefefe;
    margin: 15% auto; /* 15% from the top and centered */
    padding: 20px;
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
  }

  .close {
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;
  }

  .close:hover,
  .close:focus {
    color: black;
    text-decoration: none;
    cursor: pointer;
  }

  .modal-content button {
    margin: 10px 5px;
  }
</style>
