import fs from 'fs';

export const getInput = (inputPath: string) => {
  return fs.readFileSync(inputPath, 'utf8').trim().split('\n');
};
