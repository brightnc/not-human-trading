import React, { useState } from "react";
import axios from "axios";
import TradingViewWidget, { Themes } from "react-tradingview-widget";
import Indicator from "./Indicator";

const baseURL = "http://localhost:8000/v1/indicators"

const NewForm = () => {

  const [rsiLenght, setRsiLenght] = useState(0)

  const config = {
    headers:{
      'Content-Type': 'application/json'
    }
  }


  function createPost() {
    axios.put(baseURL, {
    rsi_config: {
        is_active: true,
       period: rsiLenght
    },
    sto_config: {
        is_active: false,
        length: 0,
        d: 0,
        k: 0
    },
    macd_config: {
        is_active: true,
        ema_fast_period: 0,
        ema_slow_period: 0,
        signal_period: 0
    },
    ema_config: {
        is_active: true,
        fast_period: 0,
        slow_period: 0
    },
    supertrend_config: {
        is_active: false,
        atr_period: 0,
        multiplier: 0
    }
}, config).then((res) => {
  console.log(res.data)
})

  }

  return (
    <div className="flex justify-between mx-8">
      
      <div className="w-full pl-5 border rounded-md ">
      <h2 className="text-2xl mb-9 text-zinc-50 font-bold  text-center">Indicator</h2>
      
        <form action="" className="" >
        <span className="mr-[5.5rem] text-zinc-50">RSI</span>
          <input
            type="text"
            id="period"
            placeholder="Lenght"
            className="border border-solid rounded-xl pl-1"
            onChange={(e) => setRsiLenght(Number(e.target.value))}
          />
          <input type="checkbox" className="ml-5 w-5 rounded-full" />
        </form>

        <form action="" className="mt-5 flex">
          <span className="mr-14 text-zinc-50">EMA</span>
          <input
            type="text"
            id="period"
            placeholder="Fast"
            className="border border-solid ml-5 rounded-xl pl-1"
          />

           <input
            type="text"
            id="period"
            placeholder="Slow"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input type="checkbox" className="ml-5 w-5 rounded-full" />
        </form>

        <form action="" className="mt-5 flex">
          <span className="mr-11 text-zinc-50">MACD</span>
          <input
            type="text"
            id="period"
            placeholder="Fast"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input
            type="text"
            id="period"
            placeholder="Slow"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input
            type="text"
            id="period"
            placeholder="Signal"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input type="checkbox" className="ml-5 w-5 rounded-full" />
        </form>

        <form action="" className="mt-5 flex">
          <span className="mr-5 text-zinc-50">Stochastic</span>
          <input
            disabled="true"
            type="text"
            id="period"
            placeholder="K Lenght 14"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input
            disabled="true"
            type="text"
            id="period"
            placeholder="D Smoothing 3"
            className="border border-solid ml-5 rounded-xl pl-1"
          />
          <input type="checkbox" className="ml-5 w-5 rounded-full" />
        </form>

        <form action="" className="mt-5 flex">
          <span className="mr-3 text-zinc-50">Supertrend</span>
          <input
            type="text"
            id="period"
            placeholder="ATR Lenght"
            className="border border-solid ml-5 rounded-xl pl-1 "
          />
          <input
            type="text"
            id="period"
            placeholder="Factor"
            className="border border-solid ml-5 rounded-xl pl-1 "
          />
          <input type="checkbox" className="ml-5 w-5 rounded-full" />
        </form>

        <button onClick={createPost} className="p-5 bg-slate-600 text-white">GO</button>
      </div>
      <TradingViewWidget symbol="BINANCE:BTCUSDT" theme={Themes.DARK} locale="en" />
    </div>
  );
};

export default NewForm;
