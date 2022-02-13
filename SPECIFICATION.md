
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

GUIDs are created (under the hood) using windows (CoCreateGuid)[https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateguid].
MS claims that they will never be the same even on different systems.


## Fields



`id`: GUID
The id of this submission.
Should be a non-random generated unique identifier


`timestamp`: string
The data this submission was created by the player. UTC time, represented in Iso8601 format.


`session_id`: GUID
A unique id that represents a single play session


`category`: string
One of "general", "bug", "performance", or "gameplay"

`mood`: uint8
A measurement from 1-4 representing the players feelings about the thing they are submitting feedback for.


`build_id`: string
The ID from unreal that identifies the build version that the player is playing

`text`: string
the text content of the submission

`state`: object
Select data about the current game and player state 
See below for the fields in this section


## State



`level_name`: string
The name of the level the player is in


`level_pos`: string "x, y, z"
The position of the player, a vector represented as a string

`playtime`: integer
The number of seconds in the game session when submitted






## Changelog

2/13/2022 Add "playtime" description & added to category tag

2/13/2022 Specify timestamp format and timezone

2/13/2022 Switch uint32 identifiers to GUIDs & add "text" field

2/10/2022 Changed player_id -> session_id

2/9/2022 Created

