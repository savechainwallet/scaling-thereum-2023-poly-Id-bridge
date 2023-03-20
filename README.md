# PolyIDBridge

This project aims to create an OAuth 2.0 adaptor for Polygon ID, enabling legacy Web 2.0 apps to obtain secure zero-knowledge proofs and passwordless authentication.

# Project Description

This project is focused on creating an OAuth 2.0 adaptor for Polygon ID, which enables legacy Web 2.0 applications to obtain secure zero-knowledge proofs and passwordless authentication. The adaptor is designed to work with off-chain verifiers and custom OpenID Connect adaptors, allowing third-party tools to use ID token OpenID Connect flows and issue JWT tokens based on JWZ tokens.

The authentication process starts when a user clicks login button on a third-party app, which redirects them to a verifier page where user can scans a QR code. The verifier then requests user data from a KYC issuer, based on the data from the JWZ token. Once the data is obtained, the verifier creates a session in Polybase and returns an ID token to the third-party app.

The third-party app then uses a client ID/secret to query the JWT token based on the session from the verifier. This enables secure authentication without requiring the user to enter a password or provide any personal information. Additionally, the use of Polybase for session caching adds an extra layer of security, ensuring that user data is stored securely and can only be accessed by authorized parties.

Overall, this project aims to provide a secure, efficient, and user-friendly method for authenticating users without the need for passwords or personal information. By leveraging the power of Polygon ID and other cutting-edge technologies, this OAuth 2.0 adaptor has the potential to revolutionize the way we think about online authentication and security.

# How it's Made

This project's backend will be written in Golang, utilizing the Polygon ID Golang SDK for secure and efficient integration with Polygon ID. Additionally, the Polybase API will be used to interact with the Polybase database for session caching, with a custom adaptor written to facilitate seamless integration.

For traditional legacy endpoints, the Gin/Gonic web framework will be utilized to handle HTTP requests and responses, while one of the JWT libraries will be used to facilitate secure authentication and authorization.

On the frontend, React will be used to create a user-friendly and intuitive interface for interacting with the application. Together, these technologies create a powerful and flexible architecture that is both secure and efficient.