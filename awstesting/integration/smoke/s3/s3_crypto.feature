# language: en
@s3 @client
Feature: S3 Integration Crypto Tests

  Scenario: Get all plaintext fixtures for symmetric masterkey aes cbc 
    When I get all fixtures for "aes_cbc" from "aws-s3-shared-tests"
    Then I decrypt each fixture against "Ruby" "version_1"
    And I compare the decrypted ciphertext to the plaintext

  Scenario: Get all plaintext fixtures for symmetric masterkey aes cbc 
    When I get all fixtures for "aes_cbc" from "aws-s3-shared-tests"
    Then I decrypt each fixture against "Java" "version_2"
    And I compare the decrypted ciphertext to the plaintext

  Scenario: Get all plaintext fixtures for symmetric masterkey aes cbc 
    When I get all fixtures for "aes_gcm" from "aws-s3-shared-tests"
    Then I decrypt each fixture against "Java" "version_2"
    And I compare the decrypted ciphertext to the plaintext
