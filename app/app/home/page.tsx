import React from 'react';

const LeetCodeStats = () => {
  const totalEasy = 738;
  const totalMedium = 1528;
  const totalHard = 636;
  const completedEasy = 500;
  const completedMedium = 500; 
  const completedHard = 250; 

  const calculatePercentage = (completed: number, total: number) => {
    return (completed / total) * 100;
  };

  return (
    //flex flex-col absolute top-40 left-500 w-80 h-96 bg-slate-400
  <div className="bg-sky-100 h-screen">
    {/* <div className="bg-slate-400 w-80 h-200 relative top-20 left-10">
      <div className="PetFloor items-baseline w-80 h-40 bg-stone-400">

      </div>
    </div> */}
    <div className="flex flex-col items-end absolute top-0 right-0 p-10">
      <div className="bg-white dark:bg-black p-4 rounded-lg shadow-md">
        <h2 className="text-xl font-semibold mb-4">LeetCode Stats</h2>
        <div className="mb-4">
          <p className="mb-2">Easy:</p>
          <div className="relative bg-green-900 bg-opacity-80 h-8 rounded-full">
            <div
              className="bg-green-600 h-8 rounded-full"
              style={{
                width: `${calculatePercentage(completedEasy, totalEasy)}%`,
              }}
            ></div>
            <p className="absolute top-1/2 left-2/4 transform -translate-y-1/2 -translate-x-1/2 text-white">
              {`${completedEasy}/${totalEasy}`}
            </p>
          </div>
        </div>
        <div className="mb-4">
          <p className="mb-2">Medium:</p>
          <div className="relative bg-gray-300 h-8 rounded-full">
            <div
              className="bg-orange-400 h-8 rounded-full"
              style={{
                width: `${calculatePercentage(completedMedium, totalMedium)}%`,
              }}
            ></div>
            <p className="absolute top-1/2 left-2/4 transform -translate-y-1/2 -translate-x-1/2 text-white">
              {`${completedMedium}/${totalMedium}`}
            </p>
          </div>
        </div>
        <div>
          <p className="mb-2">Hard:</p>
          <div className="relative bg-gray-300 h-8 rounded-full">
            <div
              className="bg-red-400 h-8 rounded-full"
              style={{
                width: `${calculatePercentage(completedHard, totalHard)}%`,
              }}
            ></div>
            <p className="absolute top-1/2 left-2/4 transform -translate-y-1/2 -translate-x-1/2 text-white">
              {`${completedHard}/${totalHard}`}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
  );
};

export default LeetCodeStats;
