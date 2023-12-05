import fs from 'fs';

export const parseInput = (inputPath: string) => {
  return fs.readFileSync(inputPath, 'utf8').trim().split('\n');
};
