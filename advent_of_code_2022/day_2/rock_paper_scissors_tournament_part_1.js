import data from './data.js';
var result = 0;

// First column of data is what the opponent will play
// A => Rock
// B => Paper
// C => Scissors

// Second column is what I should play
// X => Rock
// Y => Paper
// Z => Scissors
export const getShapeByLetter = (letter) => {
  switch (letter) {
    case "A":
    case "X":
      return "Rock";

    case "B":
    case "Y":
      return "Paper";

    case "C":
    case "Z":
      return "Scissors";
  }
};

export const myMatchResult = (myShape, oppenentShape) => {
  switch (myShape) {
    case "Rock":
      return oppenentShape === "Rock"
        ? "Draw"
        : oppenentShape === "Scissors"
        ? "Win"
        : "Lose";
    case "Paper":
      return oppenentShape === "Paper"
        ? "Draw"
        : oppenentShape === "Rock"
        ? "Win"
        : "Lose";
    case "Scissors":
      return oppenentShape === "Scissors"
        ? "Draw"
        : oppenentShape === "Paper"
        ? "Win"
        : "Lose";
  }
};

// Outcome points
// Win  => 6
// Draw => 3
// Lose => 0
export const getMatchResultPoints = (result) => {
  switch (result) {
    case "Win":
      return 6;
    case "Draw":
      return 3;
    default:
      return 0;
  }
};

// Shapes
// 1 => Rock
// 2 => Paper
// 3 => Scissors
export const getMatchShapePoints = (myShape) => {
  switch (myShape) {
    case "Rock":
      return 1;
    case "Paper":
      return 2;
    default:
      return 3;
  }
};
// Get current Match winner
// How to cacluate score for each round
// The shape + the match outcome

const matchLogic = (match) => {
  var [oppenent, me] = match;

  oppenent = getShapeByLetter(oppenent);
  me = getShapeByLetter(me);

  const matchResult = myMatchResult(me, oppenent);
  const matchScore = getMatchResultPoints(matchResult) + getMatchShapePoints(me);

  return matchScore

};

for(var i = 0; i <= data.length; i++){
  if(Array.isArray(data[i]))
    result = matchLogic(data[i]) + result;
}

console.log(result); // 9241
