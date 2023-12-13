package initializers

import (
	"fmt"
	"log"

	"github.com/gabriel-kimutai/noiter/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dn := "test.db"
	DB, err = gorm.Open(sqlite.Open(dn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}
}

var appetizers = []models.Dish{
	{
		Name:        "Crispy Calamari",
		Description: "Tender calamari rings lightly fried to perfection, served with zesty marinara sauce.",
		Price:       12.99,
		Rating:      4.5,
	},
	{
		Name:        "Spinach and Artichoke Dip",
		Description: "Creamy blend of spinach, artichokes, and melted cheese, accompanied by warm tortilla chips.",
		Price:       9.99,
		Rating:      4.2,
	},
	{
		Name:        "Stuffed Mushrooms",
		Description: "Mushroom caps filled with savory herb-infused stuffing, baked to golden perfection.",
		Price:       11.50,
		Rating:      4.3,
	},
	{
		Name:        "Bruschetta Platter",
		Description: "Toasted baguette slices topped with diced tomatoes, fresh basil, and balsamic glaze.",
		Price:       8.99,
		Rating:      4.4,
	},
	{
		Name:        "Buffalo Chicken Wings",
		Description: "Spicy buffalo wings served with cool ranch dressing and crisp celery sticks.",
		Price:       10.50,
		Rating:      4.6,
	},
	{
		Name:        "Shrimp Spring Rolls",
		Description: "Delicate spring rolls filled with succulent shrimp, crisp vegetables, and a tangy dipping sauce.",
		Price:       13.25,
		Rating:      4.2,
	},
	{
		Name:        "Caprese Skewers",
		Description: "Fresh mozzarella, cherry tomatoes, and basil drizzled with balsamic glaze on skewers.",
		Price:       9.75,
		Rating:      4.0,
	},
	{
		Name:        "Garlic Parmesan Fries",
		Description: "Crispy golden fries tossed in garlic, Parmesan, and fresh herbs, served with truffle aioli.",
		Price:       8.50,
		Rating:      4.1,
	},
	{
		Name:        "Deviled Eggs Trio",
		Description: "Classic deviled eggs with a trio of unique toppings, a delightful twist on a timeless favorite.",
		Price:       7.99,
		Rating:      4.3,
	},
	{
		Name:        "Avocado Shrimp Ceviche",
		Description: "Chilled shrimp, diced avocado, and fresh citrus combine in a refreshing ceviche.",
		Price:       12.99,
		Rating:      4.5,
	},
}

var soupsAndSalads = []models.Dish{
	{
		Name:        "Classic Caesar Salad",
		Description: "Crisp romaine lettuce tossed with Caesar dressing, Parmesan cheese, and croutons.",
		Price:       9.99,
		Rating:      4.4,
	},
	{
		Name:        "Tomato Basil Soup",
		Description: "Velvety tomato soup infused with fresh basil, served with a dollop of creamy pesto.",
		Price:       6.99,
		Rating:      4.2,
	},
	{
		Name:        "Mango Avocado Salad",
		Description: "A tropical blend of ripe mango, creamy avocado, mixed greens, and citrus vinaigrette.",
		Price:       11.50,
		Rating:      4.5,
	},
	{
		Name:        "Minestrone Soup",
		Description: "Hearty Italian vegetable soup with pasta, beans, and a savory broth.",
		Price:       7.50,
		Rating:      4.0,
	},
	{
		Name:        "Quinoa Chickpea Salad",
		Description: "Nutrient-packed quinoa and chickpeas with fresh vegetables, tossed in lemon herb dressing.",
		Price:       10.99,
		Rating:      4.3,
	},
	{
		Name:        "Lobster Bisque",
		Description: "Creamy lobster soup with a hint of sherry, garnished with lobster chunks.",
		Price:       14.99,
		Rating:      4.6,
	},
	{
		Name:        "Greek Salad",
		Description: "Crisp lettuce, tomatoes, cucumbers, olives, and feta cheese, drizzled with Greek dressing.",
		Price:       9.75,
		Rating:      4.2,
	},
	{
		Name:        "Butternut Squash Soup",
		Description: "Smooth butternut squash soup with a touch of cinnamon and nutmeg.",
		Price:       8.50,
		Rating:      4.3,
	},
	{
		Name:        "Asian Sesame Chicken Salad",
		Description: "Grilled chicken, mixed greens, and crunchy noodles, tossed in a sesame ginger dressing.",
		Price:       12.50,
		Rating:      4.1,
	},
	{
		Name:        "Roasted Red Pepper Bisque",
		Description: "Velvety bisque made with roasted red peppers, tomatoes, and a hint of cream.",
		Price:       8.99,
		Rating:      4.4,
	},
}

var mainCourses = []models.Dish{
	{
		Name:        "Grilled Salmon",
		Description: "Perfectly grilled salmon fillet with lemon herb butter, served with seasonal vegetables.",
		Price:       18.99,
		Rating:      4.7,
	},
	{
		Name:        "Chicken Alfredo",
		Description: "Creamy Alfredo sauce with tender chicken breast over fettuccine pasta.",
		Price:       15.50,
		Rating:      4.5,
	},
	{
		Name:        "Vegetarian Stir-Fry",
		Description: "Fresh mixed vegetables stir-fried in a savory sauce, served with steamed rice.",
		Price:       12.99,
		Rating:      4.4,
	},
	{
		Name:        "Beef Tenderloin Steak",
		Description: "Juicy beef tenderloin steak cooked to perfection, served with garlic mashed potatoes.",
		Price:       22.75,
		Rating:      4.8,
	},
	{
		Name:        "Shrimp Scampi",
		Description: "Lemon-garlic shrimp served over linguine pasta, a delightful seafood favorite.",
		Price:       19.99,
		Rating:      4.6,
	},
	{
		Name:        "Eggplant Parmesan",
		Description: "Slices of eggplant breaded and baked, topped with marinara sauce and melted cheese.",
		Price:       14.50,
		Rating:      4.3,
	},
	{
		Name:        "BBQ Ribs",
		Description: "Slow-cooked BBQ ribs basted in a tangy sauce, served with coleslaw and baked beans.",
		Price:       16.99,
		Rating:      4.7,
	},
	{
		Name:        "Teriyaki Chicken Bowl",
		Description: "Grilled teriyaki chicken over a bed of rice, topped with sesame seeds and green onions.",
		Price:       13.25,
		Rating:      4.4,
	},
	{
		Name:        "Mushroom Risotto",
		Description: "Creamy risotto cooked with assorted mushrooms and finished with Parmesan cheese.",
		Price:       17.50,
		Rating:      4.6,
	},
	{
		Name:        "Lobster Tail",
		Description: "Butter-poached lobster tail served with drawn butter and lemon wedges.",
		Price:       29.99,
		Rating:      4.9,
	},
}

var burgers = []models.Dish{
	{
		Name:        "Classic Cheeseburger",
		Description: "Juicy beef patty topped with melted cheese, lettuce, tomato, and pickles on a toasted bun.",
		Price:       12.50,
		Rating:      4.6,
	},
	{
		Name:        "Bacon Avocado Burger",
		Description: "Grilled burger with crispy bacon, sliced avocado, and chipotle mayo on a brioche bun.",
		Price:       13.75,
		Rating:      4.5,
	},
	{
		Name:        "Mushroom Swiss Burger",
		Description: "Beef patty topped with sautéed mushrooms and Swiss cheese, served on a pretzel bun.",
		Price:       14.25,
		Rating:      4.4,
	},
	{
		Name:        "Veggie Burger",
		Description: "Plant-based patty with lettuce, tomato, and vegan mayo on a whole grain bun.",
		Price:       11.99,
		Rating:      4.2,
	},
	{
		Name:        "BBQ Bacon Burger",
		Description: "Grilled burger topped with crispy bacon, cheddar cheese, and BBQ sauce.",
		Price:       14.50,
		Rating:      4.7,
	},
	{
		Name:        "Spicy Jalapeño Burger",
		Description: "Fiery burger with jalapeños, pepper jack cheese, and sriracha mayo on a ciabatta bun.",
		Price:       13.99,
		Rating:      4.3,
	},
	{
		Name:        "Blue Cheese Bison Burger",
		Description: "Lean bison patty topped with blue cheese, caramelized onions, and arugula on a multigrain bun.",
		Price:       15.75,
		Rating:      4.8,
	},
	{
		Name:        "Turkey Cranberry Burger",
		Description: "Ground turkey patty with cranberry sauce, stuffing, and mayo on a wheat bun.",
		Price:       12.99,
		Rating:      4.6,
	},
	{
		Name:        "Caprese Burger",
		Description: "Grilled burger with fresh mozzarella, tomato, basil, and balsamic glaze on a ciabatta bun.",
		Price:       13.25,
		Rating:      4.4,
	},
	{
		Name:        "Southwest Black Bean Burger",
		Description: "Spicy black bean patty with guacamole, salsa, and lettuce on a whole grain bun.",
		Price:       11.50,
		Rating:      4.2,
	},
}

var pizzas = []models.Dish{
	{
		Name:        "Margherita Pizza",
		Description: "Traditional pizza with tomato sauce, fresh mozzarella, basil, and a drizzle of olive oil.",
		Price:       14.99,
		Rating:      4.5,
	},
	{
		Name:        "Pepperoni Supreme Pizza",
		Description: "Classic pepperoni pizza with mushrooms, bell peppers, and black olives.",
		Price:       16.50,
		Rating:      4.6,
	},
	{
		Name:        "BBQ Chicken Pizza",
		Description: "Grilled chicken, red onions, and BBQ sauce on a thin crust, topped with cilantro.",
		Price:       17.75,
		Rating:      4.4,
	},
	{
		Name:        "Vegetarian Deluxe Pizza",
		Description: "Assorted vegetables including bell peppers, onions, olives, and mushrooms.",
		Price:       15.99,
		Rating:      4.3,
	},
	{
		Name:        "Meat Lovers' Pizza",
		Description: "A carnivore's delight with pepperoni, sausage, bacon, and ground beef.",
		Price:       18.25,
		Rating:      4.7,
	},
	{
		Name:        "Pesto Shrimp Pizza",
		Description: "Garlic shrimp, cherry tomatoes, and pesto sauce on a thin crust.",
		Price:       19.50,
		Rating:      4.8,
	},
	{
		Name:        "Hawaiian Pineapple Pizza",
		Description: "Ham, pineapple, and mozzarella cheese on a classic Hawaiian-style pizza.",
		Price:       16.99,
		Rating:      4.5,
	},
	{
		Name:        "Truffle Mushroom Pizza",
		Description: "Wild mushrooms, truffle oil, and ricotta cheese on a crispy pizza crust.",
		Price:       17.99,
		Rating:      4.6,
	},
	{
		Name:        "Buffalo Chicken Pizza",
		Description: "Spicy buffalo chicken, blue cheese, and celery on a buffalo sauce-infused crust.",
		Price:       18.50,
		Rating:      4.4,
	},
	{
		Name:        "Mediterranean Flatbread",
		Description: "Flatbread topped with hummus, feta cheese, olives, cherry tomatoes, and cucumber.",
		Price:       15.25,
		Rating:      4.2,
	},
}

var sideDishes = []models.Dish{
	{
		Name:        "Truffle Parmesan Fries",
		Description: "Crispy golden fries tossed in truffle oil and grated Parmesan, served with aioli.",
		Price:       8.99,
		Rating:      4.3,
	},
	{
		Name:        "Garlic Knots",
		Description: "Soft knots of dough brushed with garlic butter, sprinkled with herbs, and served with marinara.",
		Price:       6.50,
		Rating:      4.1,
	},
	{
		Name:        "Onion Rings",
		Description: "Beer-battered onion rings fried to a golden crisp, served with a tangy dipping sauce.",
		Price:       7.75,
		Rating:      4.2,
	},
	{
		Name:        "Sweet Potato Fries",
		Description: "Baked sweet potato fries with a touch of cinnamon and sugar, served with maple dip.",
		Price:       9.25,
		Rating:      4.4,
	},
	{
		Name:        "Caprese Salad",
		Description: "Fresh tomatoes, mozzarella, and basil drizzled with balsamic glaze.",
		Price:       9.50,
		Rating:      4.5,
	},
	{
		Name:        "Steamed Broccoli",
		Description: "Nutrient-packed steamed broccoli florets, lightly seasoned and served with lemon wedges.",
		Price:       6.99,
		Rating:      4.0,
	},
	{
		Name:        "Loaded Nachos",
		Description: "Tortilla chips smothered in cheese, guacamole, sour cream, jalapeños, and salsa.",
		Price:       10.50,
		Rating:      4.6,
	},
	{
		Name:        "Coleslaw",
		Description: "Classic coleslaw with shredded cabbage and carrots in a creamy dressing.",
		Price:       5.99,
		Rating:      4.2,
	},
	{
		Name:        "Garlic Parmesan Knots",
		Description: "Soft knots of dough brushed with garlic butter and Parmesan, served with marinara.",
		Price:       7.25,
		Rating:      4.3,
	},
	{
		Name:        "Fried Mac and Cheese Bites",
		Description: "Crispy bites of mac and cheese, served with a spicy Sriracha dipping sauce.",
		Price:       8.50,
		Rating:      4.4,
	},
}

var desserts = []models.Dish{
	{
		Name:        "Chocolate Lava Cake",
		Description: "Decadent chocolate cake with a gooey molten center, served with a scoop of vanilla ice cream.",
		Price:       9.50,
		Rating:      4.8,
	},
	{
		Name:        "New York Cheesecake",
		Description: "Creamy classic New York-style cheesecake with a graham cracker crust.",
		Price:       8.75,
		Rating:      4.6,
	},
	{
		Name:        "Tiramisu",
		Description: "Italian dessert with layers of coffee-soaked ladyfingers and mascarpone cheese, dusted with cocoa.",
		Price:       10.25,
		Rating:      4.7,
	},
	{
		Name:        "Red Velvet Cupcakes",
		Description: "Moist red velvet cupcakes with cream cheese frosting and a sprinkle of red velvet crumbs.",
		Price:       5.99,
		Rating:      4.5,
	},
	{
		Name:        "Key Lime Pie",
		Description: "Tangy key lime pie with a buttery graham cracker crust, topped with whipped cream.",
		Price:       7.50,
		Rating:      4.4,
	},
	{
		Name:        "Peach Cobbler",
		Description: "Warm peach cobbler with a buttery crust, served with a scoop of vanilla ice cream.",
		Price:       8.99,
		Rating:      4.6,
	},
	{
		Name:        "Molten Chocolate Chip Cookies",
		Description: "Warm and gooey chocolate chip cookies with a molten chocolate center, served with vanilla ice cream.",
		Price:       6.25,
		Rating:      4.3,
	},
	{
		Name:        "Banana Split",
		Description: "Classic banana split with vanilla, chocolate, and strawberry ice cream topped with whipped cream and nuts.",
		Price:       9.75,
		Rating:      4.9,
	},
	{
		Name:        "Creme Brulee",
		Description: "Silky vanilla custard with a caramelized sugar crust, a sophisticated French dessert.",
		Price:       10.99,
		Rating:      4.8,
	},
	{
		Name:        "Mixed Berry Parfait",
		Description: "Layers of mixed berries, Greek yogurt, and granola in a delightful parfait.",
		Price:       7.99,
		Rating:      4.2,
	},
}

var kidsMenu = []models.Dish{
	{
		Name:        "Cheesy Chicken Tenders",
		Description: "Crispy chicken tenders with a cheesy coating, served with a side of ketchup.",
		Price:       6.99,
		Rating:      4.2,
	},
	{
		Name:        "Mini Cheeseburgers",
		Description: "Juicy mini cheeseburgers with a soft bun, perfect for little hands.",
		Price:       5.99,
		Rating:      4.0,
	},
	{
		Name:        "Pizza Pockets",
		Description: "Mini pizza pockets filled with tomato sauce, cheese, and your choice of toppings.",
		Price:       4.50,
		Rating:      3.8,
	},
	{
		Name:        "Peanut Butter Banana Wrap",
		Description: "A delicious wrap with creamy peanut butter, sliced bananas, and a drizzle of honey.",
		Price:       3.99,
		Rating:      4.5,
	},
	{
		Name:        "Grilled Cheese Bites",
		Description: "Melted cheese sandwiched between bite-sized portions of toasted bread.",
		Price:       5.25,
		Rating:      4.3,
	},
	{
		Name:        "Fruit Kabobs",
		Description: "Colorful fruit kabobs with a variety of fresh fruits, perfect for a healthy snack.",
		Price:       4.75,
		Rating:      4.6,
	},
	{
		Name:        "Chicken Quesadilla Dippers",
		Description: "Mini chicken quesadillas served with a side of mild salsa for dipping.",
		Price:       6.50,
		Rating:      4.1,
	},
	{
		Name:        "Teddy Bear Pancakes",
		Description: "Fluffy pancake stacks shaped like teddy bears, served with syrup and fruit.",
		Price:       7.99,
		Rating:      4.4,
	},
	{
		Name:        "Spaghetti with Meatballs",
		Description: "Kid-friendly spaghetti with meatballs and marinara sauce, topped with Parmesan cheese.",
		Price:       8.25,
		Rating:      4.0,
	},
	{
		Name:        "Fruity Yogurt Parfait",
		Description: "Layers of yogurt, granola, and fresh fruit, creating a tasty and healthy parfait.",
		Price:       4.99,
		Rating:      4.7,
	},
}

var beverages = []models.Dish{
	{
		Name:        "Classic Mojito",
		Description: "A refreshing mix of muddled mint, lime juice, sugar, and rum, topped with soda.",
		Price:       8.50,
		Rating:      4.3,
	},
	{
		Name:        "Berry Blast Smoothie",
		Description: "A vibrant blend of mixed berries, yogurt, and a touch of honey for sweetness.",
		Price:       6.99,
		Rating:      4.5,
	},
	{
		Name:        "Iced Green Tea",
		Description: "Chilled green tea with a hint of citrus, served over ice for a refreshing experience.",
		Price:       4.75,
		Rating:      4.0,
	},
	{
		Name:        "Espresso Martini",
		Description: "A sophisticated mix of espresso, vodka, and coffee liqueur, shaken over ice.",
		Price:       10.25,
		Rating:      4.6,
	},
	{
		Name:        "Pineapple Coconut Cooler",
		Description: "A tropical concoction of pineapple juice, coconut cream, and a splash of sparkling water.",
		Price:       7.99,
		Rating:      4.2,
	},
	{
		Name:        "Cucumber Mint Lemonade",
		Description: "Refreshing lemonade infused with cucumber slices and fresh mint leaves.",
		Price:       5.50,
		Rating:      4.4,
	},
	{
		Name:        "Chai Latte",
		Description: "Warm and comforting chai tea mixed with steamed milk and a sprinkle of cinnamon.",
		Price:       6.25,
		Rating:      4.1,
	},
	{
		Name:        "Orange Creamsicle Smoothie",
		Description: "A creamy blend of orange juice, vanilla yogurt, and a hint of vanilla extract.",
		Price:       6.75,
		Rating:      4.7,
	},
	{
		Name:        "Cranberry Mint Spritzer",
		Description: "Cranberry juice, fresh mint, and club soda, creating a fizzy and flavorful spritzer.",
		Price:       5.99,
		Rating:      4.3,
	},
	{
		Name:        "Ginger Lemon Iced Tea",
		Description: "Iced tea infused with ginger and a splash of lemon juice for a zesty kick.",
		Price:       4.50,
		Rating:      4.5,
	},
}

var specials = []models.Dish{
	{
		Name:        "Truffle-Infused Risotto",
		Description: "Creamy risotto infused with decadent truffle oil, topped with Parmesan shavings.",
		Price:       17.99,
		Rating:      4.8,
	},
	{
		Name:        "Maple Glazed Salmon",
		Description: "Grilled salmon fillet glazed with maple syrup, served with roasted vegetables.",
		Price:       16.50,
		Rating:      4.6,
	},
	{
		Name:        "Vegetarian Stuffed Bell Peppers",
		Description: "Bell peppers filled with a flavorful mixture of quinoa, black beans, and vegetables.",
		Price:       14.99,
		Rating:      4.4,
	},
	{
		Name:        "Balsamic Strawberry Chicken",
		Description: "Grilled chicken breasts topped with a balsamic strawberry reduction, served with quinoa.",
		Price:       18.75,
		Rating:      4.7,
	},
	{
		Name:        "Miso Glazed Eggplant",
		Description: "Eggplant slices glazed with a sweet and savory miso sauce, served over rice.",
		Price:       15.25,
		Rating:      4.2,
	},
	{
		Name:        "Pesto Shrimp Linguine",
		Description: "Linguine pasta with succulent shrimp and a vibrant basil pesto sauce.",
		Price:       19.50,
		Rating:      4.5,
	},
	{
		Name:        "Blackened Chicken Tacos",
		Description: "Spicy blackened chicken, avocado, and lime slaw in soft corn tortillas.",
		Price:       16.99,
		Rating:      4.3,
	},
	{
		Name:        "Harvest Quinoa Bowl",
		Description: "A nourishing bowl with quinoa, roasted vegetables, and a tahini dressing.",
		Price:       14.50,
		Rating:      4.1,
	},
	{
		Name:        "Lemon Garlic Butter Scallops",
		Description: "Pan-seared scallops in a luscious lemon garlic butter sauce, served with orzo.",
		Price:       21.99,
		Rating:      4.9,
	},
	{
		Name:        "BBQ Pulled Pork Sliders",
		Description: "Slow-cooked pulled pork sliders with tangy barbecue sauce and coleslaw.",
		Price:       17.25,
		Rating:      4.4,
	},
}

var breakfast = []models.Dish{
	{
		Name:        "Avocado Toast with Poached Eggs",
		Description: "Sliced avocado on toasted artisan bread, topped with perfectly poached eggs.",
		Price:       10.99,
		Rating:      4.2,
	},
	{
		Name:        "Belgian Waffle Stack",
		Description: "Fluffy Belgian waffles served in a stack with maple syrup and whipped cream.",
		Price:       8.75,
		Rating:      4.0,
	},
	{
		Name:        "Eggs Florentine",
		Description: "Poached eggs on a bed of sautéed spinach, served on a toasted English muffin with hollandaise sauce.",
		Price:       12.50,
		Rating:      4.3,
	},
	{
		Name:        "Blueberry Almond Pancakes",
		Description: "Almond-infused pancakes with fresh blueberries, served with almond butter and syrup.",
		Price:       9.99,
		Rating:      4.5,
	},
	{
		Name:        "Vegetarian Breakfast Burrito",
		Description: "Scrambled eggs, black beans, peppers, and cheese wrapped in a flour tortilla, served with salsa.",
		Price:       11.25,
		Rating:      4.1,
	},
	{
		Name:        "Cinnamon Swirl French Toast",
		Description: "Thick slices of French toast with a cinnamon swirl, dusted with powdered sugar.",
		Price:       9.50,
		Rating:      4.4,
	},
	{
		Name:        "Smoked Salmon Bagel",
		Description: "Toasted bagel with cream cheese, smoked salmon, capers, and red onion slices.",
		Price:       13.99,
		Rating:      4.6,
	},
	{
		Name:        "Huevos Rancheros",
		Description: "Fried eggs on corn tortillas, topped with ranchero sauce, black beans, and guacamole.",
		Price:       10.75,
		Rating:      4.7,
	},
	{
		Name:        "Greek Yogurt Parfait",
		Description: "Layers of Greek yogurt, granola, and mixed berries for a protein-packed breakfast.",
		Price:       7.99,
		Rating:      4.2,
	},
	{
		Name:        "Sausage and Mushroom Omelette",
		Description: "Fluffy omelette filled with sausage, mushrooms, and melted cheese.",
		Price:       11.50,
		Rating:      4.8,
	},
}

var foodCategories = []models.Category{
	{
		Name:        "appetizers",
		Description: "Small dishes served before the main course to stimulate the appetite.",
		Dishes:      appetizers,
	},
	{
		Name:        "soups-and-salads",
		Description: "Various soups and salads, often offered as starters or lighter meal options.",
		Dishes:      soupsAndSalads,
	},
	{
		Name:        "main-courses",
		Description: "The main dishes of a meal, which could include meat, poultry, seafood, or vegetarian options.",
		Dishes:      mainCourses,
	},
	{
		Name:        "burgers",
		Description: "Different types of burgers, often with various toppings and side options.",
		Dishes:      burgers,
	},
	{
		Name:        "pizza",
		Description: "A variety of pizzas with different crusts, sauces, cheeses, and toppings.",
		Dishes:      pizzas,
	},
	{
		Name:        "side-dishes",
		Description: "Additional items that accompany the main course, such as fries, vegetables, or rice.",
		Dishes:      sideDishes,
	},
	{
		Name:        "desserts",
		Description: "Sweet treats to conclude the meal, including cakes, pies, ice cream, and other desserts.",
		Dishes:      desserts,
	},
	{
		Name:        "beverages",
		Description: "A selection of drinks, including soft drinks, juices, cocktails, wines, and more.",
		Dishes:      beverages,
	},
	{
		Name:        "specials",
		Description: "Unique or signature dishes recommended by the chef.",
		Dishes:      specials,
	},
	{
		Name:        "kids-menu",
		Description: "Special menu items designed for children.",
		Dishes:      kidsMenu,
	},
	{
		Name:        "breakfast-brunch",
		Description: "Dishes served during the morning hours, ranging from traditional breakfast items to brunch options.",
		Dishes:      breakfast,
	},
}

func InitializeCategories() {
	DB.Debug().Create(&foodCategories)
	fmt.Println("[INFO]: Categories Initialized")
}

func InitializeTables() {
	for i := 1; i <= 10; i++ {
		DB.Create(&models.Table{})
	}
	fmt.Println("[INFO]: Tables Initialized")
}
