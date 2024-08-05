#### Automated Programming


### Automated Code Generation:

In the context of automated code generation for the `anscombe_quartet_analysis` codebase, the original `main_test.go` file was removed, necessitating its recreation. To achieve this, a Yacc/Bison (.y) file was developed, defining the grammar and rules for generating the test file. This Yacc file specifies the structure and expected content of the test cases, including functions like `checkDataQuality`, `linearRegression`, `calculateRSquared`, and `Correlation`, as well as utility functions like `almostEqual` for floating-point comparison.

The process involved several steps:
1. **Generating the Lexer and Parser**: Using tools like `go tool yacc`, the .y file and a corresponding lexer file were parsed to generate a Go source file.
2. **Compiling the Code**: The `go build` command was used to compile the generated files and ensure they were free of syntax errors.
3. **Running the Code Generator**: `go generate` was utilized to automate the creation of boilerplate code.
4. **Executing the Generated Code**: The `go run` command was used to compile and execute the generated program, ensuring alignment with expected behavior.
5. **Running the Tests**: The `go test -v` command executed the test cases, providing verbose output for better debugging.

## Downsides and Limitations

1. **Mismatch with Original Test Cases**: The generated file does not replicate the original `main_test.go`, which contained specific test cases for critical functions. The parser focuses on arithmetic expressions, missing the original function tests.
2. **Limited Utility**: The parser does not provide practical test coverage for the codebase's functions, which is not aligned with the purpose of the original test file.
3. **Lack of Test Scenarios**: The generated file lacks comprehensive test scenarios for validating function functionality and edge cases, potentially leaving gaps in test coverage.
4. **Manual Effort Required**: Without the original test cases, developers must manually recreate them, which is time-consuming and error-prone.

## Conclusion

While `goyacc` is powerful for generating parsers, it is unsuitable for regenerating specific test cases unless explicitly defined in the grammar and actions. This situation highlights the importance of maintaining comprehensive test cases and documentation to ensure that automated tools can accurately recreate essential files when needed. Overall, it will not be recommended.

### AI Assisted Code:
GitHUB Copilot was used to write the code. Started with Import, and Anscombe Quartet data. The GitHUB Copilot predicted everything I wanted to write. Every single line of code was predicted. 

## Conclusion

Simply AMAZING. Maybe becuase I had the same code somewhere in my workspace, GitHUB Copilot predicted everything. I only typed 'func' and it wrote the rest. Compared to the other two codes, this one does not have any issues and it runs (but it does not have the correct answers in the file, in fact, the output file was empty) The other two use cases have code that does not run. Debugging was very difficult. We could start generating code with ChatGPT and then use GitHUB copilot to debug. THis migth be the fastest approach to start coding.


### AI Generated Code:
ChatGPT 4.o was used for this code generation. The chat history can be found here: https://chatgpt.com/c/01979ee9-419a-4a84-b2b0-f9296e86d99c

Observations:
1. **Input provided to ChatGPT**: Specifically requested to create Go module. Provided the Anscombe Quartet data.
2. **Error Handling**: There were number of errors when ChatGPT provided the code for this assignment. Some of these errors can be manually fixed, some of them are not. In the latter case, we have to get back to ChatGPT conversation and mention where we are getting an error. 
3. **Code Does Not Run**: Due to multiple errors that ChatGPT could not resolve, the main code does not run. 

## Conclusion

While ChatGPT was very fast to generate the codes, it was time consuming to resolve the errors it generated. I would prefer to use it to generate a simple code for each function separately. It was very fast to generate base code with some explanations. I would recommend this approach, but debugging might be difficult.
