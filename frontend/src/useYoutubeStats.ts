import { useEffect, useState } from "preact/hooks";

const SERVER_URL = "ws://localhost:8080/stats";

export function useYoutubeStats() {
  const [suscribers, setSuscribers] = useState([]);
  const socket = new WebSocket(SERVER_URL);

  useEffect(() => {
    socket.onopen = () => {
      console.log("Successfully connected");
    };

    socket.onerror = (error) => {
      console.error(`There was an error: ${JSON.stringify(error)}`);
    };

    socket.onmessage = (msg) => {
      setSuscribers(msg.data);
    };

    return (socket.onclose = () => {
      console.log("Closed");
    });
  }, []);

  return { suscribers };
}
