# Ultimate Go

Follow [Ardanlabs's Ultimate Go course](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/README.md).

## Design Philosophy

- Does your performance better? Is it your highest priority?
- Performance vs Productivity? 
- Trade-off? Everything comes with a cost
- Optimize for correctness first, then think about performance
- Code Reviews
- Integrity:
  - Be serious about writing code that reliable
  - Less code means less bugs
  - Error handling must be a part of the main code
- Readability: 
  - Not hiding the cost of the code or the decision making, the impact
- Simplicity: 
  - Hide complexity
- Performance: 
  - Compute less
  - Never guess
  - Measurements must be relevant
  - Profile
  - Test

## Syntax

- Variables
  - Type provides integrity and readability: 
    - amount of memory that we allocate and what does it present?
  - Base on computer architecture
  - Zero value
  - Short variable declaration
  - `var` is the only way to initialize a variable
  - Conversion vs casting
    - Go doesn't have casting but conversion - memory allocation
