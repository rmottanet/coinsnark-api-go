# Coin Snark API

## ctrl+s :v:

This project showcases the Coin Snark API, a currency conversion service written in Go. It provides users with a seamless way to convert monetary values between different currencies. The API is part of a larger project, which also includes a frontend page hosted on GitHub Pages for easy access to its functionalities.

While this README offers a brief overview, for a comprehensive understanding, please refer to the complete [documentation](https://rmottanet.gitbook.io/coinsnark).


## Features

Provides a list of supported currencies.
Allows currency conversion between various currency pairs.
Lightweight and easy-to-use API for seamless integration into applications.


## Usage

To use the Coin Snark API, developers need to send HTTP requests to the designated endpoints. The API responds with JSON data containing the requested information or converted currency values.


## API Endpoints:

To use the Coin Snark API locally, you can make requests to the following endpoints:

Get Supported Currency List:
```bash
curl "http://localhost:8000/api/currency"
```

Convert Currency:
```bash
curl "http://localhost:8000/api/convert?from=usd&to=brl&amount=42.75"
```

For more detailed information about the endpoints and their functionalities, refer to the central [API documentation](https://rmottanet.gitbook.io/coinsnark).

> API keys from third-party services may be required for the correct functioning of certain API functionalities.


## Contributing

Contributions to the Coin Snark API project are welcome! If you have ideas for improvements, feature requests, or bug reports, please feel free to open an issue or submit a pull request.

Thank you for considering the Coin Snark API for your currency conversion needs. If you have any questions or need further assistance, please don't hesitate to reach out. Happy coding! 🚀

<br />
<br />
<p align="center">
<a href="https://gitlab.com/rmottanet"><img src="https://img.shields.io/badge/Gitlab--_.svg?style=social&logo=gitlab" alt="GitLab"></a>
<a href="https://github.com/rmottanet"><img src="https://img.shields.io/badge/Github--_.svg?style=social&logo=github" alt="GitHub"></a>
<a href="https://instagram.com/rmottanet/"><img src="https://img.shields.io/badge/Instagram--_.svg?style=social&logo=instagram" alt="Instagram"></a>
<a href="https://www.linkedin.com/in/rmottanet/"><img src="https://img.shields.io/badge/Linkedin--_.svg?style=social&logo=linkedin" alt="Linkedin"></a>
</p>
<br />
