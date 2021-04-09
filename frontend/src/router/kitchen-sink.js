/**
 * kitchen sink imports
 * KS_ prefix is used because there was an issue with duplicate named components
 * causing infinite loops. I worked around it instead of diving deeper
 */
import KitchenSink from '../views/kitchen-sink/Index.vue'
import KS_InputText from '../views/kitchen-sink/KS_InputText.vue'
import KS_Checkbox from '../views/kitchen-sink/KS_Checkbox.vue'
import KS_Button from '../views/kitchen-sink/KS_Button.vue'
import KS_Modal from '../views/kitchen-sink/KS_Modal.vue'

// kitchen sink routes
// TODO: pull out in prod build
export const kitchenSinkRoutes = {
    path: '/kitchen-sink',
    component: KitchenSink,
    children: [
        {
            path: 'input-text',
            component: KS_InputText,
        },
        {
            path: 'checkbox',
            component: KS_Checkbox,
        },
        {
            path: 'button',
            component: KS_Button,
        },
        {
            path: 'modal',
            component: KS_Modal,
        },
    ],
}
