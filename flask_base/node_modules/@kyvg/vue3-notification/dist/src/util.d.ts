import type { NotificationItem } from './types';
interface Direction {
    x: null | string;
    y: null | string;
}
/**
  * Sequential ID generator
  */
export declare const Id: () => number;
/**
  * Splits space/tab separated string into array and cleans empty string items.
  */
export declare const split: (value: unknown) => string[];
/**
  * Cleanes and transforms string of format "x y" into object {x, y}.
  * Possible combinations:
  *   x - left, center, right
  *   y - top, bottom
  */
export declare const listToDirection: (value: string | string[]) => Direction;
export type NotificationItemWithTimer = NotificationItem & {
    timer?: number;
};
export declare class Timer {
    private start;
    private remaining;
    private notifyItem;
    private callback;
    constructor(callback: () => void, delay: number, notifItem: NotificationItemWithTimer);
    pause(): void;
    resume(): void;
}
export {};
