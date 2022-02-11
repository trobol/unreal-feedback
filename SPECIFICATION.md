
# Feedback Submission Specification

The following document defines the specification for the feedback submission document.
Feedback submissions can be created by a player while in-game and contains information about the player and game's current state (e.g. position, level, enemies killed).

A "submission" is the data inputted by a player from a single in-game form, along with the additional data from the game state and associated tracking information.
A submission is "created" when the player clicks the button to finalize the form.

Submissions are sent to a server for later processing.
The submissions are stored and transmitted as JSON with the following fields. 

Descriptions of fields in the present tense refer to the time when the submission was created.


The first section (Fields) deals with generic data that applies less to any individual game. 
The second section (State) is the data that is collect on the player.


## Fields



`id`: uint32
The id of this submission.
Should be a non-random generated unique identifier


`timestamp`: string
The data this submission was created by the player.


`session_id`: uint32
A unique id that represents a single play session


`category`: string
One of "general", "bug", "performance", or ... //TODO: missing one option (check server)

`mood`: uint8
A measurement from 0-5 representing the players feelings about the thing they are submitting feedback for.



`build_id`: string
The ID from unreal that identifies the build version that the player is playing


`state`: object
Select data about the current game and player state 
See below for the fields in this section


## State



`level_name`: string
The name of the level the player is in


`level_pos`: string "x, y, z"
The position of the player, a vector represented as a string

`playtime`: 






## Changelog

2/10/2022 Changed player_id -> session_id

2/9/2022 Created

