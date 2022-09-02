Feature: count-down
  As a counter
  In order to count accuratelly
  I want an efficient counter


  Scenario: Count down from 3 to 0
    Given I start counting from 4
    When I decrementing by 1 
    Then Print 2
    But if the numebr is 0
    Then it should print "GO"