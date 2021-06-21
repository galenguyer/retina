import { DateTime } from "luxon";
import { PureComponent } from "react";
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from "recharts";
import useSWR from "swr";
import "./App.css";
import ms from "ms";

const Graph = (props) => {
    const statuses = props.statuses,
        service = props.service;
    return (
        <div>
            <h2>
                {service} ({getUptime(statuses)}% uptime)
            </h2>
            <ResponsiveContainer height={250} width="80%">
                <LineChart data={statuses}>
                    <YAxis unit="ms" width={80} />
                    <XAxis
                        tickFormatter={formatXTick}
                        type="number"
                        domain={["dataMin", "dataMax"]}
                        tickCount={5}
                        tickLine={false}
                        axisLine={true}
                        dataKey="timestamp"
                    />
                    <Line
                        type="monotone"
                        dataKey="duration"
                        stroke="#8884d8"
                        strokeWidth={2}
                        dot={false}
                        isAnimationActive={false}
                    />
                    <Tooltip content={CustomTooltip} cursor={false} />
                </LineChart>
            </ResponsiveContainer>
        </div>
    );
};

function getUptime(statuses) {
    let up = 0,
        total = 0;
    statuses.forEach((status) => {
        total++;
        if (status.duration != null) {
            up++;
        }
    });
    return Math.round((up * 100) / total);
}

function formatXTick(ts) {
    return ms(ts - Date.now());
}

const CustomTooltip = ({ active, payload, label }) => {
    if (active) {
        return (
            <div className="custom-tooltip">
                <p className="label">{ms(Date.now() - label, { long: true })} ago</p>
                <p className="desc">{payload[0] != null ? payload[0].value + "ms" : "down"}</p>
            </div>
        );
    }
    return null;
};

export default Graph;
