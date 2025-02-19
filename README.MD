# 📝 CLI To-Do List Manager

A simple **Command-Line To-Do List Manager** built with Go 🏗️.  
It allows users to **add, list, update, and delete tasks** stored in a local JSON file.

---

## 📌 Features
✅ Add tasks with title, description, and time remaining  
✅ List all tasks (shows completion status ✅❌)  
✅ Update task details and mark tasks as completed  
✅ Delete tasks by ID  
✅ Interactive CLI menu that runs until **exit** is chosen  

---

## 🚀 Installation & Setup

### 1️⃣ **Clone the Repository**
```bash
git clone https://github.com/123456Arshia/TO_DO_CLI
cd YOUR_REPOSITORY
```

### 2️⃣ **Build the Project**
```bash
go build -o todo-cli main.go commands.go cli_command.go
```
This will generate an executable file named **`todo-cli`**.

### 3️⃣ **Run the CLI**
```bash
./todo-cli
```

---

## 📖 Usage Instructions

Once you run the program, you will see an interactive menu:

```
===== To-Do List Manager =====
1. Add Task
2. List Tasks
3. Update Task
4. Delete Task
5. Exit
Choose an option:
```

### **1️⃣ Add a Task**
```bash
# Select "1" in the menu
Enter task title: Buy groceries
Enter task description: Milk, eggs, and bread
Enter time remaining (hours): 2
```
✅ *Task added successfully!*

---

### **2️⃣ List All Tasks**
```bash
# Select "2" in the menu
===== Task List =====
1. [ ] Buy groceries (2.0 hours left)
```

---

### **3️⃣ Update a Task**
```bash
# Select "3" in the menu
Enter task ID to update: 1
Enter new title: Buy groceries and vegetables
Enter new description: Milk, eggs, bread, and carrots
Is the task completed? (yes/no): no
Enter new time remaining (hours): 1.5
```
✅ *Task updated successfully!*

---

### **4️⃣ Delete a Task**
```bash
# Select "4" in the menu
Enter task ID to delete: 1
```
✅ *Task deleted successfully!*

---

## 📂 Project Structure
```
todo-cli/
│── commands.go       # Handles task operations (reading, adding, updating, deleting)
│── cli_command.go    # Handles CLI input (asks user for input & calls functions)
│── main.go           # Main entry point, shows menu & calls CLI functions
│── Task.json         # Stores task data
```

---

## 🛠️ Technologies Used
- **Golang** – Core programming language
- **JSON** – Data storage for tasks
- **bufio.Scanner** – For interactive user input

---

## 🤝 Contributing
Feel free to submit pull requests! 🚀 If you find any issues, report them under **GitHub Issues**.

---

## 📜 License
This project is licensed under the **MIT License**. You are free to modify and distribute it.

---

