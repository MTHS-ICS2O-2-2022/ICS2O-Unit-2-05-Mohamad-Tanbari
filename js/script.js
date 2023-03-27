// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mr. Coxall
// Created on: Sep 2020
// This file contains the JS functions for index.html

/*
* This function calcualtes how much money you made post tax
*/
function calculate () {
  // debugging statements
  console.log('button clicked')

  // recieve input from user
  const hourlyWage = parseInt(document.getElementById('hourly-wage').value)
  const hoursWorked = parseInt(document.getElementById('hours-worked').value)
  const TAX_RATE = 0.18 // 18%

  // calcuations
  const grossPay = hourlyWage * hoursWorked
  const netPay = grossPay - (grossPay * TAX_RATE)

  // output result into answerbox
  document.getElementById('answer').innerHTML = 'Your net pay is $' + netPay.toFixed(2)
}
