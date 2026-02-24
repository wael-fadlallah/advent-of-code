import fs from "fs";

const data = fs.readFileSync("puzzle_input.txt", "utf-8").split("\n");

const getLetterPriorityNumber = (letter) => {
  // use charCodeAt to return the ascii value of the letter and subtract 64 to make the values start from 1
  const n = letter.charCodeAt() - 64;
  // now we check if the letter value is more than 33 in the ascii table that means it's a small letter to lets map the value between 1-26
  if (n >= 33) return n - 32;

  // If it's not grater than 32 that means it's a captial letter to let's map the value between 27-52
  return n + 26;
};

const getCompartments = (str) => {
  return [
    str.substr(0, str.length / 2),
    str.substr(str.length / 2, str.length - 1),
  ];
};

let answer = 0;
data.map(line => {
  // WE iterate through every line and get the two compartments then we get the duplicate letters  
  const [compartment_1, compartment_2] = getCompartments(line);
  let shared_item = compartment_1.split('').filter(letter => compartment_2.includes(letter));
  shared_item = [...new Set(shared_item)]

  if(shared_item.length > 0) answer += getLetterPriorityNumber(shared_item.join());
});
console.log({part_1: answer});


// Part 2 
let answer_2 = 0;

for(let i = 0; i < data.length - 3; i+=3){

  const [line1, line2, line3] = [data[i], data[i+1], data[i+2]];

  let shared_item_1_2 = line1.split('').filter(letter => line2.includes(letter));
  let shared_item_2_3 = shared_item_1_2.filter(letter => line3.includes(letter));

  let shared_item = [...new Set(shared_item_2_3)];
  answer_2 += getLetterPriorityNumber(shared_item.join(''))
}

console.log({part_2: answer_2});
