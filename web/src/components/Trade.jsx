import React, { useState } from "react";
import axios from "axios";
import TradingViewWidget, { Themes } from "react-tradingview-widget";
import Run from "./Run";
import Form from "react-bootstrap/Form";
import Logs from "./logs";

const baseURL = "http://localhost:8000/v1/bots";

const Trade = () => {
  const [isActive, setIsActive] = useState(false);
  const [selectRSI, setSelectRSI] = useState(false);
  const [selectEMA, setSelectEMA] = useState(false);
  const [selectMACD, setSelectMACD] = useState(false);
  const [selectSTO, setSelectSTO] = useState(false);
  const [selectSUPER, setSelectSUPER] = useState(false);

  const [configs, setConfigs] = useState({
    symbol: "",
    qty: "",
  });

  const [tf, setTf] = useState("");

  const [rsiLenght, setRsiLenght] = useState(0);
  const [ema, setEma] = useState({
    fast: 0,
    slow: 0,
  });
  const [macd, setMacd] = useState({
    fast: 0,
    slow: 0,
    signal: 0,
  });

  const [supertrend, setSupertrend] = useState({
    atr: 0,
    multiplier: 0,
  });

  const config = {
    headers: {
      "Content-Type": "application/json",
    },
  };

  var data = {
    rsi_config: {
      is_active: selectRSI,
      period: rsiLenght,
    },
    sto_config: {
      is_active: selectSTO,
      length: 14,
      d: 3,
      k: 1,
    },
    macd_config: {
      is_active: selectMACD,
      ema_fast_period: macd.fast,
      ema_slow_period: macd.slow,
      signal_period: macd.signal,
    },
    ema_config: {
      is_active: selectEMA,
      fast_period: ema.fast,
      slow_period: ema.slow,
    },
    supertrend_config: {
      is_active: selectSUPER,
      atr_period: supertrend.atr,
      multiplier: supertrend.multiplier,
    },
    order_config: {
      sym: configs.symbol,
      qty: configs.qty,
    },
    timeframe: tf,
  };

  function handleStart() {
    axios.post(baseURL + "/start", data, config).then((res) => {
      console.log(res.data);
    });
    setIsActive(true);
  }

  function handleStop() {
    axios.post(baseURL + "/stop", config).then((res) => {
      console.log(res.data);
    });
    setIsActive(false);
  }

  const changeHandlerConfigs = (e) => {
    setConfigs({ ...configs, [e.target.name]: e.target.value });
  };

  const changeHandlerEMA = (e) => {
    setEma({ ...ema, [e.target.name]: Number(e.target.value) });
  };

  const changeHandlerMACD = (e) => {
    setMacd({ ...macd, [e.target.name]: Number(e.target.value) });
  };
  const changeHandlerSuperTrend = (e) => {
    setSupertrend({ ...supertrend, [e.target.name]: Number(e.target.value) });
  };

  return (
    <>
      <div className="flex ">
        <form
          action=""
          className=" flex items-center justify-center mb-3 font-main"
        >
          <span className="mx-[5rem] text-[#848e9c]">Symbol</span>
          <input
            name="symbol"
            type="text"
            id=""
            className="border border-solid rounded-xl pl-1"
            placeholder="BTCUSDT"
            onChange={changeHandlerConfigs}
          />

          <span className="mx-[5rem] text-[#848e9c]">Qty</span>
          <input
            name="qty"
            type="text"
            id=""
            className="border border-solid rounded-xl pl-1"
            pattern=""
            onChange={changeHandlerConfigs}
          />
        </form>
      </div>

      <div className="w-auto h-[600px] mx-8 bg-transparent bg-black">
        <div className="w-1/2 float-left pl-5 border rounded-md h-1/2 p-9 ">
          <div className="flex w-full">
            <h2 className="text-2xl mt-2 mb-6 text-[#F3c535] font-bold  text-center">
              Indicator
            </h2>

            <div className="mt-2 ml-11">
              <Form.Select
                aria-label="Default select example"
                onChange={(e) => setTf(e.target.value)}
              >
                <option>Timeframe</option>
                <option value="1m">1 m</option>
                <option value="3m">3 m</option>
                <option value="5m">5 m</option>
                <option value="15m">15 m</option>
                <option value="30m">30 m</option>
                <option value="1h">1 hour</option>
                <option value="2h">2 hours</option>
                <option value="4h">4 hours</option>
                <option value="6h">6 hours</option>
                <option value="8h">8 hours</option>
                <option value="12h">12 hours</option>
                <option value="1d">1 day</option>
                <option value="3d">3 days</option>
                <option value="1w">1 week</option>
                <option value="1mth">1 month</option>
              </Form.Select>
              {/* <span className="mr-5 text-[#848e9c]">Time Frame</span>
            <input
            name="timeframe"
              type="text"
              id=""
              className="border border-solid rounded-xl pl-3 w-16 h-8"
              pattern=""
              onChange={changeHandlerConfigs}
            /> */}
            </div>
          </div>
          <div className="w-full">
            <form action="" className="mt-2 ">
              <span className="mr-[5.5rem] text-[#848e9c]">RSI</span>
              <input
                type="number"
                id="rsiLenght"
                placeholder="Lenght"
                className="border border-solid rounded-xl pl-1 py-0 text-sm"
                onChange={(e) => setRsiLenght(Number(e.target.value))}
                pattern=""
              />
              <input
                type="checkbox"
                checked={selectRSI}
                className="ml-5 w-5 rounded-full"
                onChange={() => setSelectRSI(!selectRSI)}
              />
            </form>

            <form action="" className="mt-3 flex">
              <span className="mr-14 text-[#848e9c]">EMA</span>
              <input
                name="fast"
                type="number"
                id="emaF"
                placeholder="Fast"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm"
                onChange={changeHandlerEMA}
              />

              <input
                name="slow"
                type="number"
                id="emaS"
                placeholder="Slow"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm"
                onChange={changeHandlerEMA}
              />
              <input
                type="checkbox"
                checked={selectEMA}
                className="ml-5 w-5 rounded-full"
                onChange={() => setSelectEMA(!selectEMA)}
              />
            </form>

            <form action="" className="mt-3 flex">
              <span className="mr-11 text-[#848e9c]">MACD</span>
              <input
                name="fast"
                type="number"
                id="macdF"
                placeholder="Fast"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm w-40"
                onChange={changeHandlerMACD}
              />
              <input
                name="slow"
                type="number"
                id="macdS"
                placeholder="Slow"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm w-40"
                onChange={changeHandlerMACD}
              />
              <input
                name="signal"
                type="number"
                id="macdSig"
                placeholder="Signal"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm w-40"
                onChange={changeHandlerMACD}
              />
              <input
                type="checkbox"
                checked={selectMACD}
                className="ml-5 w-5 rounded-full"
                onChange={() => setSelectMACD(!selectMACD)}
              />
            </form>

            <form action="" className="mt-3 flex">
              <span className="mr-2 text-[#848e9c]">Stochastic</span>
              <input
                disabled="true"
                type="number"
                placeholder="K Lenght 14"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm"
              />
              <input
                disabled="true"
                type="number"
                placeholder="D Smoothing 3"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm"
              />
              <input
                type="checkbox"
                checked={selectSTO}
                className="ml-5 w-5 rounded-full"
                onChange={() => setSelectSTO(!selectSTO)}
              />
            </form>

            <form action="" className="mt-3 flex">
              <span className="mr-1 text-[#848e9c]">Supertrend</span>
              <input
                name="atr"
                type="number"
                id="atr"
                placeholder="ATR Lenght"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm "
                onChange={changeHandlerSuperTrend}
              />
              <input
                name="multiplier"
                type="number"
                id="multiplier"
                placeholder="Factor"
                className="border border-solid ml-5 rounded-xl pl-1 py-0 text-sm "
                onChange={changeHandlerSuperTrend}
              />
              <input
                type="checkbox"
                checked={selectSUPER}
                className="ml-5 w-5 rounded-full"
                onChange={() => setSelectSUPER(!selectSUPER)}
              />
            </form>
          </div>
        </div>

        <div className="w-1/2 h-full float-right">
          <TradingViewWidget
            symbol="BINANCE:BTCUSDT"
            theme={Themes.DARK}
            locale="en"
            autosize
          />
        </div>
        <div className="w-1/2 h-1/2 float-left relative">
          <Logs />
          {/* <div className="bg-white  float-left w-[49%] h-[95%]  absolute left-0 bottom-0 overflow-y-scroll ">LEFT</div>
            <div className="bg-white float-left w-[49%]  h-[95%] absolute right-0 bottom-0 overflow-y-scroll">RIGHT</div> */}
        </div>
      </div>
      <div className="flex items-center absolute left-96 gap-2 mt-3">
        <Run
          isActive={isActive}
          title="START"
          color="green"
          onClick={handleStart}
        />
        <Run
          isActive={!isActive}
          title="STOP"
          color="red"
          onClick={handleStop}
        />
      </div>
    </>
  );
};

export default Trade;
