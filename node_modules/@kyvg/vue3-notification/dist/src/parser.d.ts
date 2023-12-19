export type ValueType = {
    type: string;
    value: number | string;
};
export declare const parse: (value: number | string) => ValueType;
export default parse;
