<h1>Booking Application</h1>

This code is a simple CLI booking application for a Go conference. It allows users to book tickets for the conference and sends them a confirmation email with the tickets.


<h3>Concepts learnt:</h3>

This project served as a valuable learning experience, allowing me to grasp Go concepts through hands-on development. Here's a breakdown of the key areas explored:

**Package Management:** I gained experience with importing functionalities from external packages like _fmt_ for printing and _time_ for time manipulation.

**Data Structures:** Building the _UserData_ struct helped me understand how to structure complex data containing user information (name, email, tickets).

**Control Flow:** The _for_ loop and _if_ statements provided practice in controlling program flow based on user input and validation rules.

**Functions:** Defining and utilizing various functions like _greetUsers_, _getUserInput_, and _bookTickets_ solidified my understanding of modularity and code organization.

**Concurrency (Introduction):** Launching the _sendTickets_ function as a goroutine introduced me to the concept of concurrent programming in Go.

**Slices and Pointers (Introductory):** Working with slices like _bookings_ and using the address-of operator (_&_) during user input helped me grasp these fundamental Go concepts.


<h3>Project:</h3>

Here's a breakdown of the functionalities of this project:

**Tracks Conference Information:** The code stores details like the conference name, total number of tickets, and remaining tickets.

**User Interaction:** It greets users, displays available tickets, and prompts them to enter details like name, email, and the number of tickets they want to booking.

**Input Validation:** The code validates user-provided information like names and email format, and ensures they are requesting a valid number of tickets (less than or equal to remaining tickets).

**Booking Logic:** If the user's information is valid, the code books the tickets, updates remaining ticket count, and stores the user's information.

**Confirmation and Notifications:** It sends a confirmation message to the user with the number of tickets booked and their email address. It also prints a message indicating how many tickets are remaining.

**Asynchronous Ticket Delivery (Simulation):** The code simulates sending an email with the tickets by using a goroutine and sleeps for 10 seconds.
Overall, this code demonstrates a basic conference booking application with user input validation, data storage, and asynchronous email delivery simulation.


This project provided a well-rounded introduction to Go, allowing me to learn by doing and solidify my understanding of the language's core functionalities.

<h3>Output:</h3>

This is how the code would work incase the user is inputing accurate information.

![image](https://github.com/user-attachments/assets/b64fb820-06e3-4c1c-ab25-36bb6b859f49)

![image](https://github.com/user-attachments/assets/7e565fd5-eb36-4ade-90b3-6e2623e7ee4e)

![image](https://github.com/user-attachments/assets/f41dadd5-b02f-4801-bf76-67fd0259b323)

![image](https://github.com/user-attachments/assets/57310663-88be-417b-bdbc-1c5af4ac67a1)

Incase of any errors in input from the user, the terminal will generate output showing what error has ocurred:

![4](https://github.com/user-attachments/assets/04d3f2be-3672-4855-b31b-8d3afc769a40)

