# language: en
@codecommit @client
Feature: Amazon CodeCommit

  Scenario: Making a basic request
    When I call the "ListRepositories" API
    Then the value at "Repositories" should be a list

  Scenario: Error handling
    When I attempt to call the "GetRepository" API with:
    | RepositoryName | bogus-repo |
    Then I expect the response error code to be "RepositoryDoesNotExistException"
