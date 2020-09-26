import * as React from "react";
import { Checkbox, Label, NumericInput } from "@blueprintjs/core";
import { TimePicker, TimePrecision } from "@blueprintjs/datetime";
import { StrategyPanel } from "../../OrderTicket";

export interface Props {
    children?: React.ReactNode
}

export interface State {
}

export interface VwapParamsState {

    defaultStartTime: Date;
    defaultEndTime: Date;
    useDefaultBuckets: boolean;
}


export default class VwapParamsPanel extends React.Component<Props, State> implements StrategyPanel {

    public state = {
        defaultStartTime: new Date(),
        defaultEndTime: new Date(new Date().getTime() + 60000),
        buckets: 10,
        setBuckets: false
    };

    startTime: Date;
    endTime: Date;
    buckets: number;

    constructor(props: Props) {
        super(props)

        this.startTime = this.state.defaultStartTime
        this.endTime = this.state.defaultEndTime
        this.buckets = this.state.buckets
    }

    getDestination(): string {
        return "XVWAP"
    }


    getParamsString(): string {

        if (this.state.setBuckets) {
            return JSON.stringify(new VwapParameters(Math.floor(this.startTime.getTime() / 1000), Math.floor(this.endTime.getTime() / 1000), this.buckets))
        } else {
            return JSON.stringify(new VwapParameters(Math.floor(this.startTime.getTime() / 1000), Math.floor(this.endTime.getTime() / 1000)))
        }
    }

    render() {

        let defaultBuckets;
        if (this.state.setBuckets) {
            defaultBuckets = <div>
                <Label htmlFor="input-b" style={{ paddingRight: 5, paddingLeft: 25 }}>Buckets</Label>
                <NumericInput defaultValue={10} min={1} style={{ maxWidth: 60 }}
                    onValueChange={(valueAsNumber: number, valueAsString: string, inputElement: HTMLInputElement | null) => { this.buckets = valueAsNumber }} />
            </div>
        }

        return (
            <div>
            <div style={{ display: 'flex', flexDirection: 'row', paddingTop: 25, alignItems: "center" }}>
                <Label htmlFor="input-b" style={{ paddingRight: 5 }}>Start</Label>
                <TimePicker showArrowButtons={true} defaultValue={this.state.defaultStartTime} precision={TimePrecision.MINUTE} onChange={(time: Date) => { this.startTime = time }} ></TimePicker>
                <Label htmlFor="input-b" style={{ paddingRight: 5, paddingLeft: 25 }}>  End </Label>
                <TimePicker showArrowButtons={true} defaultValue={this.state.defaultEndTime} precision={TimePrecision.MINUTE} onChange={(time: Date) => { this.endTime = time }}></TimePicker>
                </div>
                <div>
                <Checkbox style={{ paddingRight: 5, paddingLeft: 25 }} checked={this.state.setBuckets} onChange={() => { this.onSetBucketsChanged() }}>Set Buckets</Checkbox>
                {defaultBuckets}
                </div>
            
            </div>
        )
    }

    private onSetBucketsChanged() {
        this.setState({ ...this.state, setBuckets: !this.state.setBuckets });
    }

}





class VwapParameters {
    utcStartTimeSecs: number;
    utcEndTimeSecs: number;
    buckets?: number;

    constructor(utcStartTimeSecs: number,
        utcEndTimeSecs: number,
        buckets?: number) {
        this.utcStartTimeSecs = utcStartTimeSecs;
        this.utcEndTimeSecs = utcEndTimeSecs;
        this.buckets = buckets;
    }
}


