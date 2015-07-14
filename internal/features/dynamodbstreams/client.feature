# language: en
@dynamodbstreams
Feature: Amazon DynamoDB Streams

  I want to use Amazon DynamoDB Streams

  Scenario: Listing streams
    When I call the "ListStreams" API
    Then the value at "Streams" should be a list

  Scenario: Error handling
    When I attempt to call the "DescribeStream" API with:
    | StreamArn | fake-stream |
    Then I expect the response error code to be "ValidationException"
