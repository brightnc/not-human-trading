import React, { useEffect, useRef, useState } from "react";
const URL = "ws://localhost:8000/ws/v1";
const isBrowser = typeof window !== "undefined";
let tradignReports = [];
const Logs = () => {
  const [ws, setWebSocket] = useState(null);
  const [isOpen, setIsOpen] = useState(false);

  const [msg, setMsg] = useState("");
  const [time, setTime] = useState("");
  const [type, setType] = useState("");
  const [messages, setMessages] = useState({});
  const [msgArr, setMsgArr] = useState([]);

  useEffect(() => {
    handleWebSocket();
  }, [ws]);

  // Only set up the websocket once
  // if (!clientRef.current) {
  //   const webSocket = new Websocket(URL, 'echo-protocol');
  //   clientRef.current = webSocket;

  //   window.webSocket = webSocket;

  //   webSocket.onerror = (e) => console.error(e);

  const addMessage = (v) => {
    setMsgArr([v, ...msgArr]);
  };

  if (ws) {
    ws.onopen = () => {
      console.log("ws opened");
      setIsOpen(true);
    };

    ws.onclose = () => {
      setIsOpen(false);
      setTimeout(() => {
        reconnect();
      }, 1000);
      console.log("ws closed");
    };
  }

  // // // Setting this will trigger a re-run of the effect,
  // // // cleaning up the current websocket, but not setting
  // // // up a new one right away
  // // setWaitingToReconnect(true);

  // // This will trigger another re-run, and because it is false,
  // // the socket will be set up again
  // setTimeout(() => setWaitingToReconnect(null), 5000);

  const reconnect = () => {
    if (!isOpen) {
      setWebSocket(new WebSocket(URL));
      setIsOpen(true);
    }
  };

  const handleWebSocket = () => {
    if (isBrowser) {
      if (ws) {
        ws.onmessage = (message) => {
          const parsedData = JSON.parse(message.data);
          // setMsgArr(oldArray => [...oldArray, message.data.toString()]);
          if (parsedData.type === "TRADING_REPORT") {
            console.log("got TRADING_REPORT type !!! ");
            tradignReports.push({
              message: parsedData.message,
              time: parsedData.time,
              type: parsedData.type,
            });
          }
          setMessages(parsedData);
          setTime(parsedData.time);
          setMsg(parsedData.message);
          setType(parsedData.type);
        };
      } else {
        setWebSocket(new WebSocket(URL));
      }
    }
  };

  //   const feedItems = feeds.map((f) => (
  //     {

  // //  <li key={f.time}>{d.message}</li>);
  //   }
  //   )
  const yearFormat = time.slice(0, 10);
  const timerFormat = time.slice(11, 19);

  const TradingReportList = () => {
    return (
      <div className=" text-sm text-white overflow-y-scroll w-64 h-48">
        {tradignReports.map((f) => {
          return <li key={f.time}>{f.message}</li>;
        })}
      </div>
    );
  };

  // const WebsocketSection = (message) => {
  // return (
  //   <>
  //     <div className="bg-transparent float-left w-[49%] h-[95%]  absolute left-0 bottom-0 ">
  //       <h1 className="text-yellow-300">
  //         ws {isOpen ? "Connected" : "Disconnected"}
  //       </h1>
  //       {ws && <p className="text-yellow-300">Reconnecting momentarily...</p>}
  //       <p className="text-sm text-white">
  //         🕑 Time : {timerFormat} {yearFormat}
  //       </p>
  //       <p className=" text-sm text-white">
  //         {" "}
  //         🔴 Status : {message.type === "FEED" ? message.message : "ABCDS"}
  //       </p>
  //     </div>

  //     <div className="bg-transparent float-left w-[49%]  h-[95%] absolute right-0 bottom-0">
  //       <h1 className=" text-sm text-yellow-300"> 🔴 Ordered</h1>
  //       {message.type === "FEED" ?
  //         <div className=" text-sm text-white"><TradingReportList message={message}/></div>
  //        : (
  //         "SSSSSS"
  //       )}
  //     </div>
  //   </>
  // );

  // }

  {
  }

  return (
    <>
      <div className="bg-transparent float-left w-[49%] h-[95%]  absolute left-0 bottom-0 ">
        <h1 className="text-yellow-300">
          Server {isOpen ? "Connected" : "Disconnected"}
        </h1>
        {/* {ws && <p className="text-yellow-300">Reconnecting momentarily...</p>} */}
        <p className="text-sm text-white">
          🕑 Time : {timerFormat} {yearFormat}
        </p>
        <p className=" text-sm text-white">
          {" "}
          🔴 Status :{" "}
          {messages.type === "FEED"
            ? messages.message
            : "Waiting for Connection..."}
        </p>
      </div>

      <div className="bg-transparent float-left w-[49%]  h-[95%] absolute right-0 bottom-0">
        <h1 className=" text-sm text-yellow-300"> 🔴 Ordered</h1>
        {tradignReports.length > 0 ? (
          <TradingReportList />
        ) : (
          <p className=" text-sm text-white">Waiting for Order...</p>
        )}
      </div>
    </>
  );
};

export default Logs;
