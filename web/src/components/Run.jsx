import React, { useState } from "react";
import { PlayCircleOutline, PowerSettingsNew } from "@mui/icons-material";
const Run = ({ title, onClick, color, isActive }) => {
  return (
    <div>
      <button
        disabled={isActive}
        onClick={onClick}
        className={`p-4 bg-${color}-600 text-white rounded-full ${
          isActive ? "cursor-not-allowed" : ""
        }`}
      >
        {title == "START" ? (
          <PlayCircleOutline fontSize="large" />
        ) : (
          <PowerSettingsNew fontSize="large" />
        )}
        {title}
      </button>
    </div>
  );
};

export default Run;
