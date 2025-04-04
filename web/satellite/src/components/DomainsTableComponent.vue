// Copyright (C) 2024 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <v-text-field
        v-model="search"
        label="Search"
        :prepend-inner-icon="Search"
        single-line
        variant="solo-filled"
        flat
        hide-details
        clearable
        density="comfortable"
        class="mb-5"
    />

    <v-data-table
        :sort-by="sortBy"
        :headers="headers"
        :items="domains"
        :search="search"
    >
        <template #item.name="{ item }">
            <v-list-item class="font-weight-bold pl-0">
                {{ item.name }}
            </v-list-item>
        </template>
        <template #item.createdAt="{ item }">
            <span class="text-no-wrap">
                {{ Time.formattedDate(item.createdAt) }}
            </span>
        </template>
    </v-data-table>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import {
    VDataTable,
    VTextField,
    VListItem,
} from 'vuetify/components';
import { Search } from 'lucide-vue-next';

import { Domain } from '@/types/domains';
import { useDomainsStore } from '@/store/modules/domainsStore';
import { Time } from '@/utils/time';
import { DataTableHeader } from '@/types/common';

type SortItem = {
    key: keyof Domain;
    order: boolean | 'asc' | 'desc';
};

const domainsStore = useDomainsStore();

const search = ref<string>('');
const sortBy = ref<SortItem[] | undefined>([{ key: 'createdAt', order: 'asc' }]);

const headers: DataTableHeader[] = [
    {
        title: 'Name',
        align: 'start',
        key: 'name',
    },
    { title: 'Date Created', key: 'createdAt' },
];

const domains = computed<Domain[]>(() => domainsStore.state.domains);
</script>
