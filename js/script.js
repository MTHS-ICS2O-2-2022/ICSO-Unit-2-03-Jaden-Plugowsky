// Copyright (c) 2023 Jaden Plugowsky All rights reserved
//
// Created by: Jaden Plugowsky
// Created on: March 2023
// This file contains the JS functions for index.html

function buttonOneClicked() {
  //Input through Textfields
  const streetName = document.getElementById("street-name").value
  const streetNumber = parseInt(document.getElementById("street-number").value)

  //Output to text
  document.getElementById("answer").innerHTML =
    "Your address is: " + streetNumber + " " + streetName + "."
}
