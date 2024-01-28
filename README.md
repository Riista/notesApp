# NotesTool

Notes Tool is a command-line tool for managing collections of short single-line notes. Each collection has its own database, stored as a plain text file with the same name as the collection. Notes are stored in separate rows within this file. The data persists between runs of the tool, ensuring that users can access and manage their notes across multiple sessions.

## Usage

Users can 
1. Create a collection of notes
2. View existing notes from the collection
3. Add new notes to the collection
4. Remove existing notes from the collection
5. Exit the program

To use Notes Tool user must execute the following command on terminal:

$./notestool [COLLECTION_NAME]

If wrong number of arguments is providen, short help message is shown:

$ ./notestool
Usage: ./notestool [COLLECTION_NAME]

When given a right arguments, the tool greets user and the following menu is shown:

1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

User can then select a command of their choice and manage the current collection.