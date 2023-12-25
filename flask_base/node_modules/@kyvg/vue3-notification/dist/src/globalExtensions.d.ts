import { notify } from './index';
declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        $notify: typeof notify;
    }
}
