import { DateTime } from "luxon";
import { PureComponent } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from "recharts";
import useSWR from "swr";
import "./App.css";
import ms from "ms";
import Graph from "./Graph";

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

    var sorted = {};
    var pingtimes = {};
    rawData.forEach((point) => {
        if (sorted[point.servicename] == null) {
            sorted[point.servicename] = [];
            pingtimes[point.servicename] = [];
        }
        sorted[point.servicename].push(point);
        pingtimes[point.servicename].push({
            timestamp: Number(new Date(point.timestamp)),
            duration: point.success ? point.duration : null,
        });
    });

    pingtimes = sortObj(pingtimes);

    return (
        <div className="App">
            <h1>Retina</h1>
            {Object.keys(pingtimes).map((service) => {
                return <Graph service={service} statuses={pingtimes[service]} />;
            })}
        </div>
    );
}

function sortObj(obj) {
    return Object.keys(obj)
        .sort()
        .reduce(function (result, key) {
            result[key] = obj[key];
            return result;
        }, {});
}

export default App;
