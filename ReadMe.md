# Expense Tracker

Sample solution for the [expense-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/DevSatyamCollab/expense-tracker.git
cd expense-tracker
```

Run the following command to build and run the project:

```bash
go build -o expense-tracker
rm -r expense-tracker/

./expense-tracker -h # To see the list of available commands

# To add a expense
./expense-tracker -add --desc "Breakfast" --amount 20 --categ "food"

# To update a expense
./expense-tracker upd 1 --desc "Lunch" --amount 30 --categ "food"

# To delete a expense
./expense-tracker del 1

# total summary
./expense-tracker -sum

# summary of categories
./expense-tracker -sum --cat

# summary of the month
./expense-tracker -sum --month 1

# To list all expenses
./expense-tracker -list

# List all expense of the month
./expense-tracker -list --month 1

# list all expense of the category
./expense-tracker -list --categ "food"
```
