<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->

<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <!-- <a href="https://github.com/lxkedinh/dsearch">
    <img src="images/logo.png" alt="Logo" width="80" height="80">
  </a> -->

<h3 align="center">dsearch</h3>

  <p align="center">
    Fuzzy file searcher CLI app written in Go!
    <br />
    <a href="https://github.com/lxkedinh/dsearch"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!-- <a href="https://github.com/lxkedinh/dsearch">View Demo</a> -->
    <!-- · -->
    <a href="https://github.com/lxkedinh/dsearch/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/lxkedinh/dsearch/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <!-- <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul> -->
    </li>
    <li><a href="#usage">Usage</a>
      <ul>
        <li><a href="#flags">Flags</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <!-- <li><a href="#contributing">Contributing</a></li> -->
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <!-- <li><a href="#acknowledgments">Acknowledgments</a></li> -->
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

<!-- [![Product Name Screen Shot][product-screenshot]](https://example.com) -->

This is a small passion project of mine I made to learn Go! It's a small CLI application that
fuzzy searches directories for files.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

- [![Golang][Go]][Go-url]
- [Cobra](https://cobra.dev/)
- [Spinner](https://github.com/briandowns/spinner)
- [normalize](https://github.com/avito-tech/normalize)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

You can get started with dsearch by downloading the latest release on the right sidebar for your
platform.

<!-- ### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

- npm
  ```sh
  npm install npm@latest -g
  ``` -->

<!-- 1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/lxkedinh/dsearch.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = "ENTER YOUR API";
   ``` -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

dsearch only requires one argument, your search query, such as:

```sh
dsearch foo
```

to fuzzy search for "foo".

### Flags

- `-d` or `--dir`: specify the directory to start fuzzy searching in
  - defaults to current directory where command was called from
- `-t` or `--threshold`: the threshold to determine strictness of fuzzy matching
  - allowed range (lowest to highest strictness): 0.0 to 1.0
  - defaults to 0.5

<!-- _For more examples, please refer to the [Documentation](https://example.com)_ -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [ ] Look into searching in parallel, learn channels and goroutines for concurrency

See the [open issues](https://github.com/lxkedinh/dsearch/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

<!-- ## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Luke Dinh - lukexdinh@gmail.com

[![Gmail][gmail-shield]][gmail-url]
[![LinkedIn][linkedin-shield]][linkedin-url]
[![GitHub][github-shield]][github-url]

Project Link: [https://github.com/lxkedinh/dsearch](https://github.com/lxkedinh/dsearch)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

<!-- ## Acknowledgments

- []()
- []()
- []()

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/lxkedinh/dsearch.svg?style=for-the-badge
[contributors-url]: https://github.com/lxkedinh/dsearch/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/lxkedinh/dsearch.svg?style=for-the-badge
[forks-url]: https://github.com/lxkedinh/dsearch/network/members
[stars-shield]: https://img.shields.io/github/stars/lxkedinh/dsearch.svg?style=for-the-badge
[stars-url]: https://github.com/lxkedinh/dsearch/stargazers
[issues-shield]: https://img.shields.io/github/issues/lxkedinh/dsearch.svg?style=for-the-badge
[issues-url]: https://github.com/lxkedinh/dsearch/issues
[license-shield]: https://img.shields.io/github/license/lxkedinh/dsearch.svg?style=for-the-badge
[license-url]: https://github.com/lxkedinh/dsearch/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-0A66C2.svg?style=for-the-badge&logo=linkedin
[linkedin-url]: https://linkedin.com/in/lxkedinh
[gmail-shield]: https://img.shields.io/badge/Gmail-EA4335?style=for-the-badge&logo=Gmail&logoColor=white&logoSize=auto
[gmail-url]: mailto:lukexdinh@gmail.com
[github-shield]: https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=GitHub&logoColor=white&logoSize=auto
[github-url]: https://github.com/lxkedinh
[product-screenshot]: images/screenshot.png
[Go]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white&logoSize=auto
[Go-url]: https://go.dev/
