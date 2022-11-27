import React from "react";
import Header from "./components/Header";
import Trade from "./components/Trade";
import TradingViewWidget, { Themes } from "react-tradingview-widget";
import Form from "./components/Form";
import NewForm from "./components/NewForm";
import Run from "./components/Run";

function App() {
  return (
    <>
      <Header />
      <Trade />
    </>
  );
}

export default App;
