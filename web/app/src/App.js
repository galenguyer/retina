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
        pingtimes[point.servicename].push({ timestamp: Number(new Date(point.timestamp)), duration: point.success ? point.duration : null })
    });

    const service = "google";

    return (
        <div className="App">
            <h1>bad graph time</h1>
            <div>
            <h2>{service} ({getUptime(pingtimes[service])}% uptime)</h2>
            <ResponsiveContainer height={250} width="80%" >
                <LineChart data={pingtimes[service]}>
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
        </div>
    );
}

function getUptime(statuses) {
    let up = 0, total = 0;
    statuses.forEach(status => {
        total++;
        if (status.duration != null) {
            up++;
        }
    })
    return Math.round(up*100/total);
}

function formatXTick(ts) {
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
