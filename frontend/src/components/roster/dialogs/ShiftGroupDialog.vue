<script setup lang="ts">
import { ref, watch } from 'vue';
import type { RosterShift, ShiftGroup } from '@gewis/grooster-backend-ts';
import { dialogPt, fieldSelectPt, readableDisabled } from '@/theme/pt';

const props = defineProps<{
    shift: RosterShift | null;
    shiftGroups: ShiftGroup[];
}>();
const visible = defineModel<boolean>('visible', { required: true });
const emit = defineEmits<{ submit: [shiftGroupId: number] }>();

const groupId = ref<number | null>(null);
watch(visible, (open) => {
    if (open) groupId.value = props.shift?.shiftGroupId ?? null;
});

function submit() {
    if (groupId.value == null) return;
    emit('submit', groupId.value);
    visible.value = false;
}
</script>

<template>
    <Dialog v-model:visible="visible" header="Shift group" modal :pt="dialogPt">
        <form id="shift-group-form" class="flex flex-col gap-1.5" @submit.prevent="submit">
            <label class="text-[13px] font-semibold text-surface-700 dark:text-surface-100" for="shift-group-select">
                Group for {{ shift?.name }}
            </label>
            <Select
                v-model="groupId"
                class="h-[42px] w-full items-center"
                input-id="shift-group-select"
                option-label="name"
                option-value="id"
                :options="shiftGroups"
                placeholder="Choose a group"
                :pt="fieldSelectPt"
            />
            <p v-if="shiftGroups.length === 0" class="text-sm text-surface-600 dark:text-surface-200">
                This organ has no shift groups yet. Create them on the templates page.
            </p>
        </form>
        <template #footer>
            <Button class="h-10 text-sm" label="Cancel" outlined severity="secondary" @click="visible = false" />
            <Button
                :class="['h-10 text-sm', readableDisabled]"
                :disabled="groupId == null"
                form="shift-group-form"
                label="Save"
                type="submit"
            />
        </template>
    </Dialog>
</template>
