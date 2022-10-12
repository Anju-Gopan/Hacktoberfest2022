/* A simple quiz app using NodeJS */

const readline = require('readline');
const chalk = require('chalk');

/**const emoji = require('node-emoji');
const tada = emoji.get('pizza');**/

const rl = readline.createInterface({
 input: process.stdin,
 output: process.stdout
});

/**const r1 = readline.createInterface({
 input: process.stdin,
 output: process.stdout
});**/

console.log(chalk.bgGreen.bold.blue('HEY! WELCOME TO GK QUIZ\n'));
console.log(chalk.bgWhite.black.bold("<---Instructions--->\n*Enter the option corresponding to each question as answer\n*For correct answer +2 points !!\n*For wrong answer -1 "));

/**r1.question(chalk.bold.bgWhite.black("Please Enter your name"), (name) => {
  console.log(chalk.red.bold(`Hello  ${name}!\n`));
});**/ 


var qbank = [
    {
      question: `
    1. Which is the National Animal of India ? 
    (a) Lion
    (b) Tiger
    (c) Peacock
    (d) Elephant\n`,
      answer: "b"
    },

    {
      question: `
    2. Who among the following wrote Sanskrit grammar?
(a) Kalidasa
(b) Charak
(c) Panini
(d) Aryabhatt\n`,
      answer: "c"
    },

    {
      question: `
    3. The metal whose salts are sensitive to light is?
(a) Zinc
(b) Silver
(c) Copper
(d) Aluminum\n`,
      answer: "b"
    },

    {
      question: `
    4. Chambal river is a part of –
(a) Sabarmati basin
(b) Ganga basin
(c) Narmada basin
(d) Godavari basin\n`,
      answer: "c"
    },

    {
      question: `
    5. FFC stands for
(a) Federation of Football Council
(b) Film Finance Corporation
(c) Foreign Finance Corporation
(d) None of the above\n`,
      answer: "b"
    },

    {
      question: `
    6. The United Nations Organization has its Headquarters at
(a) Bali
(b) Hague
(c) New York, USA
(d) Washington DC\n`,
      answer: "c"
    }
  ];

var index = 0;
var score = 0;

function check()
{
    rl.question(qbank[index].question, (answer) => {
    console.log(`Your answer is  ${answer}`);

    if(answer.toLowerCase() == qbank[index].answer){
      console.log(chalk.green("Correct!"));
      index++;
      score = score+2;
      status();
    }
    else{
      console.log(chalk.red("Wrong!"));
      index++;
      if(score>0){
      score--;
      }
      status();
    }
  });
}

function status()
{
  if(qBankEnd()){
    console.log(chalk.redBright.bgWhite.bold("Thank you for playing!\n"));
    console.log(chalk.green.bgRed.bold("Your Score is: "+score));
    rl.close();
  }

  else
  {
    rl.question("Press c to continue or Press e to exit---->", (choice) => {
      //console.log(`Your choice is  ${choice}`);
      if (choice.toLowerCase() != "e"){
        check();

      }

      else{
        console.log(chalk.redBright.bgWhite.bold("Thank you for playing!\n"));
        console.log(chalk.green.bgRed.bold("Your Score is: "+score));
        rl.close();
      }
    });
  }
}
    function qBankEnd()
    {
      if(qbank.length == index){
        return true;
      }
    }

status();
Footer
