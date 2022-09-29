//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.0, generator: @autorest/go@4.0.0-preview.42)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package dataprotectiondatasourceplugin

import "encoding/json"

func unmarshalBackupCriteriaClassification(rawMsg json.RawMessage) (BackupCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupCriteriaClassification
	switch m["objectType"] {
	case "ScheduleBasedBackupCriteria":
		b = &ScheduleBasedBackupCriteria{}
	default:
		b = &BackupCriteria{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBackupCriteriaClassificationArray(rawMsg json.RawMessage) ([]BackupCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BackupCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBackupCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBackupDatasourceParametersClassification(rawMsg json.RawMessage) (BackupDatasourceParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupDatasourceParametersClassification
	switch m["objectType"] {
	case "BlobBackupDatasourceParameters":
		b = &BlobBackupDatasourceParameters{}
	default:
		b = &BackupDatasourceParameters{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBackupDatasourceParametersClassificationArray(rawMsg json.RawMessage) ([]BackupDatasourceParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]BackupDatasourceParametersClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalBackupDatasourceParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalBackupParametersClassification(rawMsg json.RawMessage) (BackupParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BackupParametersClassification
	switch m["objectType"] {
	case "AzureBackupParams":
		b = &AzureBackupParams{}
	case "AzureBackupParamsForPlugin":
		b = &AzureBackupParamsForPlugin{}
	default:
		b = &BackupParameters{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalBaseResourcePropertiesClassification(rawMsg json.RawMessage) (BaseResourcePropertiesClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b BaseResourcePropertiesClassification
	switch m["objectType"] {
	case "VmwareVMProperties":
		b = &VmwareVMProperties{}
	default:
		b = &BaseResourceProperties{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataStoreParametersClassification(rawMsg json.RawMessage) (DataStoreParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b DataStoreParametersClassification
	switch m["objectType"] {
	case "AzureOperationalStoreParameters":
		b = &AzureOperationalStoreParameters{}
	default:
		b = &DataStoreParameters{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalDataStoreParametersClassificationArray(rawMsg json.RawMessage) ([]DataStoreParametersClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]DataStoreParametersClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalDataStoreParametersClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalItemLevelRestoreCriteriaClassification(rawMsg json.RawMessage) (ItemLevelRestoreCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b ItemLevelRestoreCriteriaClassification
	switch m["objectType"] {
	case "ItemPathBasedRestoreCriteria":
		b = &ItemPathBasedRestoreCriteria{}
	case "KubernetesPVRestoreCriteria":
		b = &KubernetesPVRestoreCriteria{}
	case "KubernetesStorageClassRestoreCriteria":
		b = &KubernetesStorageClassRestoreCriteria{}
	case "RangeBasedItemLevelRestoreCriteria":
		b = &RangeBasedItemLevelRestoreCriteria{}
	default:
		b = &ItemLevelRestoreCriteria{}
	}
	return b, json.Unmarshal(rawMsg, b)
}

func unmarshalItemLevelRestoreCriteriaClassificationArray(rawMsg json.RawMessage) ([]ItemLevelRestoreCriteriaClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []json.RawMessage
	if err := json.Unmarshal(rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]ItemLevelRestoreCriteriaClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalItemLevelRestoreCriteriaClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return fArray, nil
}

func unmarshalTriggerContextClassification(rawMsg json.RawMessage) (TriggerContextClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(rawMsg, &m); err != nil {
		return nil, err
	}
	var b TriggerContextClassification
	switch m["objectType"] {
	case "AdhocBasedTriggerContext":
		b = &AdhocBasedTriggerContext{}
	case "ScheduleBasedTriggerContext":
		b = &ScheduleBasedTriggerContext{}
	default:
		b = &TriggerContext{}
	}
	return b, json.Unmarshal(rawMsg, b)
}
