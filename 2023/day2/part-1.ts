import { getInput } from '../helpers/getInput';

const sumPossibleGames = (inputPath: string) => {
  const amountPossible = {
    red: 12,
    blue: 14,
    green: 13,
  };

  const input = getInput(inputPath);
};

console.log(sumPossibleGames('./example1-input.txt'));
