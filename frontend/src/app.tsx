import { useYoutubeStats } from "./useYoutubeStats";

export function App() {
  const { suscribers } = useYoutubeStats();

  return (
    <>
      <h1 className="text-3xl font-bold underline">Youtbe Monitoring</h1>
      <p>Suscribers: {JSON.stringify(suscribers.length)}</p>
    </>
  );
}
