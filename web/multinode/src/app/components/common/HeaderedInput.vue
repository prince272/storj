// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

<template>
    <div class="input-container">
        <div v-if="!isOptional" class="label-container">
            <div class="label-container__main">
                <div v-if="error" class="label-container__main__error-icon-container">
                    <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <rect width="20" height="20" rx="10" fill="#EB5757" />
                        <path d="M10.0012 11.7364C10.612 11.7364 11.1117 11.204 11.1117 10.5532V5.81218C11.1117 5.75302 11.108 5.68991 11.1006 5.63074C11.0192 5.06672 10.5565 4.62891 10.0012 4.62891C9.39037 4.62891 8.89062 5.16138 8.89062 5.81218V10.5492C8.89062 11.204 9.39037 11.7364 10.0012 11.7364Z" fill="white" />
                        <path d="M10.0001 12.8906C9.13977 12.8906 8.44531 13.5851 8.44531 14.4454C8.44531 15.3057 9.13977 16.0002 10.0001 16.0002C10.8604 16.0002 11.5548 15.3057 11.5548 14.4454C11.5583 13.5851 10.8638 12.8906 10.0001 12.8906Z" fill="white" />
                    </svg>
                </div>
                <h3 v-if="!error" class="label-container__main__label">{{ label }}</h3>
                <h3 v-if="!error" class="label-container__main__label add-label">{{ additionalLabel }}</h3>
                <h3 v-if="error" class="label-container__main__error">{{ error }}</h3>
            </div>
            <h3 v-if="isLimitShown" class="label-container__limit">{{ currentLimit }}/{{ maxSymbols }}</h3>
        </div>
        <div v-if="isOptional" class="optional-label-container">
            <h3 class="label-container__label">{{ label }}</h3>
            <h4 class="optional-label-container__optional">Optional</h4>
        </div>
        <textarea
            v-if="isMultiline"
            :id="label"
            v-model="value"
            class="headered-textarea"
            :placeholder="placeholder"
            :style="style.inputStyle"
            :rows="5"
            :cols="40"
            wrap="hard"
            autocomplete="off"
            @input="onInput"
            @change="onInput"
            @paste.prevent="onPaste"
        />
        <input
            v-if="!isMultiline"
            :id="label"
            v-model="value"
            class="headered-input"
            :placeholder="placeholder"
            :type="[isPassword ? 'password': 'text']"
            autocomplete="off"
            :style="style.inputStyle"
            @input="onInput"
            @change="onInput"
            @paste.prevent="onPaste"
        >
    </div>
</template>

<script lang="ts">
import { Component, Prop } from 'vue-property-decorator';

import HeaderlessInput from './HeaderlessInput.vue';

// Custom input component with labeled header.
// @vue/component
@Component
export default class HeaderedInput extends HeaderlessInput {
    @Prop({ default: '' })
    private readonly initValue: string;
    @Prop({ default: '' })
    private readonly additionalLabel: string;
    @Prop({ default: 0 })
    private readonly currentLimit: number;
    @Prop({ default: false })
    private readonly isOptional: boolean;
    @Prop({ default: false })
    private readonly isLimitShown: boolean;
    @Prop({ default: false })
    private readonly isMultiline: boolean;

    public value: string;

    public created() {
        this.value = this.initValue;
    }
}
</script>

<style scoped lang="scss">
    .input-container {
        display: flex;
        flex-direction: column;
        align-items: flex-start;
        margin-top: 10px;
        font-family: 'font_regular', sans-serif;
    }

    .label-container {
        width: 100%;
        display: flex;
        justify-content: space-between;
        align-items: center;

        &__main {
            display: flex;
            justify-content: flex-start;
            align-items: center;
            margin-bottom: 8px;

            &__label {
                font-family: 'font_regular', sans-serif;
                font-size: 16px;
                line-height: 21px;
                color: var(--v-text-base);
            }

            &__error {
                font-size: 16px;
                line-height: 21px;
                color: var(--c-error);
                margin-left: 10px;
            }

            &__error-icon-container {
                width: 20px;
                height: 20px;
                max-width: 20px;
                max-height: 20px;
                display: flex;
            }
        }

        &__limit {
            font-size: 16px;
            line-height: 21px;
            color: rgb(56 75 101 / 40%);
        }
    }

    .optional-label-container {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        width: 100%;

        &__optional {
            font-size: 16px;
            line-height: 21px;
            color: var(--v-text-base);
        }
    }

    .headered-textarea {
        padding: 15px 22px;
        text-indent: 0;
        line-height: 26px;
    }

    .headered-input,
    .headered-textarea {
        font-size: 16px;
        line-height: 21px;
        resize: none;
        height: 48px;
        width: 100%;
        text-indent: 20px;
        padding-right: 20px;
        outline: none;
        box-shadow: none;
        font-family: 'font_regular', sans-serif;
        border: 1px solid var(--v-border-base);
        border-radius: var(--br-input);
        color: var(--v-text-base);
        caret-color: var(--c-primary);
        box-sizing: border-box;

        &::placeholder {
            color: var(--c-placeholder);
        }

        &:hover {
            border: 2px solid var(--v-header-base);
        }

        &:focus,
        &:active {
            border: 2px solid var(--v-primary-base);
        }
    }

    .add-label {
        margin-left: 5px;
        color: rgb(56 75 101 / 40%);
    }
</style>
