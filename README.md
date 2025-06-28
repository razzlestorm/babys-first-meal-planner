# babys-first-meal-planner
# TODO
* Create grid template for meal/snack times and days of the week.
    * Create func to print things out neatly
* Create randomization rules (like unique foods every day, not too much of one food per week, consecutively, etc.)
* Create customization rules (for enabling/disabling foods)
* Create an Orchestration with a controller for instances of the meal planning calendar for each user.
    * Should update cal in real time for changes users make
    * Should be in goroutines that begin and end when user logs in/out
* ~~Decide on actual DB and use it.~~
    * Create API for db, allow users to add foods (normalize and/or suggest a food if it's already in the db)
* Create frontend interface that calls various endpoints
