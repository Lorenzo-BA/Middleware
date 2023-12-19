import { NotificationsOptions } from './types';
export declare const notify: {
    (args: NotificationsOptions | string): void;
    close(id: unknown): void;
};
export declare const useNotification: () => {
    notify: {
        (args: NotificationsOptions | string): void;
        close(id: unknown): void;
    };
};
