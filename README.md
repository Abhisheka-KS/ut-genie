# ut-genie

**ut-genie** is an automated unit test case generation tool designed to streamline your testing workflow. Leveraging the power of the GPT-4 base model—enhanced and fine-tuned with our custom intelligence—ut-genie automatically generates unit test cases for all functions changed in a Pull Request (PR). Once generated, the test cases are added as a separate commit to the same PR via a GitHub Actions workflow, ensuring your codebase remains robust and well-tested.

## Table of Contents

- [Features](#features)
- [How It Works](#how-it-works)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Features

- **Automated Test Generation:** Automatically generates unit test cases for every function changed in a PR.
- **AI-Powered:** Utilizes the GPT-4 base model, fine-tuned with our domain-specific intelligence, for accurate and relevant test generation.
- **Seamless Integration:** Integrates with GitHub Actions to run as part of your CI/CD pipeline.
- **Separate Commit:** Automatically adds the generated tests to the same PR as a separate commit, keeping your commit history clean and organized.
- **Customizable:** Easily configurable to fit your project’s structure and testing framework.

## How It Works

1. **PR Detection:** When a pull request is opened or updated, the GitHub Actions workflow is triggered.
2. **Change Analysis:** The tool analyzes the diff to identify all functions that have been modified.
3. **Test Generation:** For each changed function, ut-genie uses the GPT-4 base model to generate relevant unit test cases.
4. **Commit Creation:** The generated test cases are automatically committed as a separate commit to the same PR.
5. **Review & Merge:** Developers review the generated tests, make any necessary adjustments, and merge the PR.

![alt text](https://github.com/Abhisheka-KS/ut-genie/master/image.png?raw=true)

## Getting Started

### Prerequisites

- **Git:** Ensure Git is installed on your machine.
- **GitHub Repository:** A repository where you can integrate the workflow.
- **GitHub Actions:** Access to GitHub Actions to run the CI/CD workflow.

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/pure-px/ut-genie/
   cd ut-genie
