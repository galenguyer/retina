import { DateTime } from "luxon";
import { PureComponent } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import useSWR from "swr";
import "./App.css";
import ms from "ms";

function App() {
    let { data: rawData, error: error } = useSWR("/api/v1/hour");

    if (error) {
        console.error(error);
        return (
            <div className="App">
                <h1>Retina</h1>
                <h2>An error occurred</h2>
            </div>
        );
    }
    if (!rawData)
        return (
            <div className="App">
                <h1>Retina</h1>
                <h2>Loading latest data...</h2>
            </div>
        );

    var sorted = {}
    var pingtimes = {}
    rawData.forEach(point => {
        if (sorted[point.servicename] == null) {
            sorted[point.servicename] = []
            pingtimes[point.servicename] = []
        }
        sorted[point.servicename].push(point)
        pingtimes[point.servicename].push({ timestamp: Number(new Date(point.timestamp)), duration: point.duration > 0 ? point.duration : null })
    });
    console.log(sorted)
    console.log(pingtimes["google"])


    return (
        <div className="App">
            <h1>bad graph time</h1>
            <ResponsiveContainer height={250} width="80%">
                <LineChart data={pingtimes["google"]}>
                    <YAxis
                        unit="ms"
                        width={80}
                    />
                    <XAxis
                        tickFormatter={formatXTick}
                        type="number"
                        domain={['dataMin', 'dataMax']} 
                        tickCount = {5}
                        tickLine={false}
                        axisLine={true}
                        dataKey="timestamp"
                    />
                    <Line type="monotone" dataKey="duration" stroke="#8884d8" strokeWidth={2} dot={false} isAnimationActive={false} />
                    <Tooltip content={CustomTooltip} cursor={false} />
                </LineChart>
            </ResponsiveContainer>

        </div>
    );
}

function formatXTick(ts) {
    console.log(ts)
    return ms(ts - Date.now())
}

const CustomTooltip = ({ active, payload, label }) => {
    if (active) {
        return (
            <div className="custom-tooltip">
                <p className="label">
                    {ms(label- Date.now(), { long: true })}
                </p>
                <p className="desc">{payload[0] != null ? payload[0].value + "ms" : "down"}</p>
            </div>
        );
    }
    return null;
}

export default App;
