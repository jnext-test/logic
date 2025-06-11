## Adding a Node to a Singly Linked List (Append Operation)

Here's a clear breakdown of the steps to add a new node to the **end** of a singly linked list:

1.  **Create the New Node**:
    First, you need to **create an instance of your node class**. This new node will contain the `data` you want to store and its `next` pointer will initially be `NULL` (or `None`), as it will become the last node.

2.  **Check if the List is Empty**:
    Before doing anything else, you must **check if the list's `head` is `NULL`**. This tells you if the list currently has any nodes.

3.  **If the List is Empty**:
    If the `head` is indeed `NULL`, it means the list is empty. In this case, your `new_node` simply becomes the **`head` of the list**.

4.  **If the List is Not Empty**:
    If the list already contains data (i.e., `head` is not `NULL`), you cannot just add the new node directly. You need to **traverse the list from the `head` to find the very last node**.

5.  **Traverse and Add**:
    Start with a temporary pointer (let's call it `current`) initialized to the `head`. Keep moving `current` to `current.next` until `current.next` becomes `NULL`. This indicates that `current` is now pointing to the last node in the list. Once you've found the last node, simply **set its `next` pointer to your `new_node`**.


## Deleting a Node to a Singly Linked List (Delete Operation)

1. **If list is empty**
  check if list is empty then return
2. **If list is not empty**
  if list is not empty then we have traverse the List
3 **find the elemet**
  find the element to remove
4 **important thing how the flow work after finding the correct element**
  like when are traversing 1 -> 2 -> 3 -> 4 -> 5 so we want to delete 3
  we have to check one step future so we can skip the correct element.
  how? so lets say we are at 2 and in if check we are checking next.next === 3
  so in this case we will skip current.next = current.next.next
5 **in final we will return**
  oho problame is solved
