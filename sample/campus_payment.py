import json
import os
from datetime import datetime

# File to store user data (JSON format)
USER_DATA_FILE = "users.json"
# File to store transaction history
TRANSACTION_FILE = "transactions.txt"

def load_users():
    """Load user data from JSON file. Create file if missing."""
    if not os.path.exists(USER_DATA_FILE):
        return {}
    with open(USER_DATA_FILE, "r") as f:
        return json.load(f)

def save_users(users):
    """Save user data to JSON file."""
    with open(USER_DATA_FILE, "w") as f:
        json.dump(users, f, indent=2)

def record_transaction(sender, receiver, amount, action):
    """Log transactions to a text file."""
    timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    log_entry = f"{timestamp} | {action} | From: {sender} | To: {receiver} | Amount: ${amount:.2f}\n"
    with open(TRANSACTION_FILE, "a") as f:
        f.write(log_entry)

def register_user():
    """Register a new user with a starting balance of $100."""
    users = load_users()
    username = input("Enter username: ").strip()

    if username in users:
        print("Username already exists!")
        return

    password = input("Enter password: ").strip()
    users[username] = {"password": password, "balance": 100.00}
    save_users(users)
    print("Registration successful! You received a $100 sign-up bonus.")

def login():
    """Authenticate user and return username if successful."""
    users = load_users()
    username = input("Username: ").strip()
    password = input("Password: ").strip()

    if username in users and users[username]["password"] == password:
        print(f"Welcome, {username}!")
        return username
    else:
        print("Invalid credentials.")
        return None

def user_menu(username):
    """Logged-in user menu."""
    users = load_users()

    while True:
        print("\n===== CampusPay CLI =====")
        print(f"User: {username} | Balance: ${users[username]['balance']:.2f}")
        print("1. Check Balance")
        print("2. Send Money")
        print("3. Request Money")
        print("4. View Transactions")
        print("5. Logout")

        choice = input("> ").strip()

        if choice == "1":  # Check Balance
            print(f"Your balance: ${users[username]['balance']:.2f}")

        elif choice == "2":  # Send Money
            recipient = input("Recipient's username: ").strip()
            amount = float(input("Amount: $").strip())

            # Validate transaction
            if recipient not in users:
                print("Recipient not found.")
            elif amount <= 0:
                print("Amount must be positive.")
            elif users[username]["balance"] < amount:
                print("Insufficient funds.")
            else:
                # Update balances
                users[username]["balance"] -= amount
                users[recipient]["balance"] += amount
                save_users(users)
                record_transaction(username, recipient, amount, "SEND")
                print("Payment sent!")

        elif choice == "3":  # Request Money
            target = input("Request from (username): ").strip()
            amount = float(input("Amount: $").strip())

            if target not in users:
                print("User not found.")
            elif amount <= 0:
                print("Amount must be positive.")
            else:
                record_transaction(target, username, amount, "REQUEST")
                print(f"Request for ${amount:.2f} sent to {target}.")

        elif choice == "4":  # View Transactions
            try:
                with open(TRANSACTION_FILE, "r") as f:
                    print("\n--- Transaction History ---")
                    print(f.read())
            except FileNotFoundError:
                print("No transactions yet.")

        elif choice == "5":  # Logout
            print("Logged out.")
            break

        else:
            print("Invalid choice.")

def main():
    """Main CLI loop."""
    while True:
        print("\n===== CampusPay CLI =====")
        print("1. Register")
        print("2. Login")
        print("3. Exit")
        choice = input("> ").strip()

        if choice == "1":
            register_user()
        elif choice == "2":
            user = login()
            if user:
                user_menu(user)
        elif choice == "3":
            print("Goodbye!")
            break
        else:
            print("Invalid choice.")

if __name__ == "__main__":
    main()