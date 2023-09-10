The project is expected to be completed in 5 milestones. These can be found [here](https://github.com/cliche-niche/CS455/milestones). The average time required for each milestone is estimated to be ~28 hours. It is planned to be implemented in Go. <br>
Milestone-wise (and task/issue-wise) breakup of the expected time is provided below. These estimates have been made taking into consideration the previous experience of the team members of working in Go, and other heuristics, such as taking [#1](https://github.com/cliche-niche/CS455/issues/1) to be the reference. The estimated time mentioned corresponding to each issue here is <b>per person</b>; the total time a milestone is estimated to take is mentioned at the end.

+ [M1: Formation of a basic text space](https://github.com/cliche-niche/CS455/milestone/3) <br>
This milestone is concerned with providing a basic interface to interact with the application, supporting file handling, etc. The exact issues included in this milestone are:
    + [#1 Handle create, save and open existing text files](https://github.com/cliche-niche/CS455/issues/1): Estimated to take 2 hours since these options are expected to be provided directly by the language.
    + [#2 Kind of data structure to use in the text buffer](https://github.com/cliche-niche/CS455/issues/2): Estimated to take a lot of time (~3 hours) since it includes a lot of research (for the purpose of ideation and feasibility) and its implementation.
    + [#5 Basic Navigation](https://github.com/cliche-niche/CS455/issues/5)
    + [#7 Basic Editing](https://github.com/cliche-niche/CS455/issues/7)
    + [#18 Creating a terminal window for the Editor](https://github.com/cliche-niche/CS455/issues/18): Estimated to take 2 hours (including [#5](https://github.com/cliche-niche/CS455/issues/5) and [#7](https://github.com/cliche-niche/CS455/issues/7)). <br>

  Estimated total human hours required for this milestone: 4 * (2 + 3 + 2) = 28
+ [M2: Text Placement and Navigation](https://github.com/cliche-niche/CS455/milestone/5) <br>
This milestone is concerned with providing basic navigation functionalities. The exact issues included in this milestone are:
    + [#6 Advanced Navigation using Hotkeys](https://github.com/cliche-niche/CS455/issues/6): Estimated to take 3 hours.
    + [#8 Display line numbers on the side of the text area](https://github.com/cliche-niche/CS455/issues/8): Estimated to take 1.5 hours.
    + [#9 Add ability to scroll using the mouse wheel in the buffer](https://github.com/cliche-niche/CS455/issues/9): Estimated to take 1 hour. <br>

  Estimated total human hours required for this milestone: 4 * (3 + 1.5 + 1) = 22
+ [M3: Advanced Editing](https://github.com/cliche-niche/CS455/milestone/6) <br> 
This milestone is concerned with providing advanced editing features like multiple cursors, moving text blocks, etc. The exact issues included in this milestone are:
    + [#12 Cutting/Copying lines](https://github.com/cliche-niche/CS455/issues/12)
    + [#15 Moving text blocks](https://github.com/cliche-niche/CS455/issues/15): Estimated to take 5 hours along with [#12](https://github.com/cliche-niche/CS455/issues/12). The two issues have been clubbed because it is anticipated that their implementation will be similar, since their semantics are also similar.
    + [#13 Multiple cursors](https://github.com/cliche-niche/CS455/issues/13)
    + [#16 Edit multiple occurrences](https://github.com/cliche-niche/CS455/issues/16): Estimated to take 5 hours along with [#13](https://github.com/cliche-niche/CS455/issues/13). The two issues have been clubbed because it is anticipated that their implementation will be similar, since their semantics are also similar.<br>
    
  Estimated total human hours required for this milestone: 4 * (5 + 5) = 40
+ [M4: Text Formatting](https://github.com/cliche-niche/CS455/milestone/4) <br>
This milestone is concerned with providing basic text formatting like bold, and providing markdown rendering/preview. The exact issues included in this milestone are:
    + [#3 Basic highlighting of keywords, numbers, etc.](https://github.com/cliche-niche/CS455/issues/3): Estimated to take 1.5 hours since a reference for "coloring" is already provided. It is expected that implementing other (similar) functionalities like bold, italics, etc. should not take a lot of time.
    + [#4 Markdown rendering](https://github.com/cliche-niche/CS455/issues/4): Estimated to take 6 hours since relatively big task.<br>
    
  Estimated total human hours required for this milestone: 4 * (1.5 + 6) = 30
+ [M5: Other Features](https://github.com/cliche-niche/CS455/milestone/7) <br>
This milestone is concerned with extending the project to include some more useful features like autosave, find and replace, etc. The exact issues included in this milestone are:
    + [#10 Find and Replace](https://github.com/cliche-niche/CS455/issues/10): Estimated to take 1 hour since some similar functionalities <i>should have been</i> implemented in a prior milestone (for eg. [M3/#16](https://github.com/cliche-niche/CS455/issues/16) is expected to be similar).
    + [#14 Autosave after certain intervals Primary](https://github.com/cliche-niche/CS455/issues/14): Estimated to take 1 hour.
    + [#17 Extend "Find and Replace" using RegEx](https://github.com/cliche-niche/CS455/issues/17): Estimated to take 1 hour since [M5/#10](https://github.com/cliche-niche/CS455/issues/10) <i>ought to have been</i> implemented before this.<br>
    
  Estimated total human hours required for this milestone: 4 * (1 + 1 + 1) = 12

Estimated total human hours required for the whole project: 28 + 22 + 40 + 30 + 12 = 132