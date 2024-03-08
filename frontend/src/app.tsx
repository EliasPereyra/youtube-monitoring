import { useState } from "preact/hooks";

export function App() {
  const [suscribers, setSuscribers] = useState(0);

  return (
    <>
      <h1 className="text-3xl font-bold underline">Youtbe Monitoring</h1>
      <p>Suscribers: {suscribers}</p>
    </>
  );
}
