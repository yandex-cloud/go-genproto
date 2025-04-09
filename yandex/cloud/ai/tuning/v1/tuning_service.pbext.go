// Code generated by protoc-gen-goext. DO NOT EDIT.

package fomo

func (m *ListTuningsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListTuningsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTuningsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTuningsRequest) SetStatus(v TuningTask_Status) {
	m.Status = v
}

func (m *ListTuningsResponse) SetTuningTasks(v []*TuningTask) {
	m.TuningTasks = v
}

func (m *ListTuningsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *DescribeTuningRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *DescribeTuningResponse) SetTuningTask(v *TuningTask) {
	m.TuningTask = v
}

func (m *CancelTuningRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *CancelTuningResponse) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *TuningResponse) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *TuningResponse) SetStatus(v TuningTask_Status) {
	m.Status = v
}

func (m *TuningResponse) SetTargetModelUri(v string) {
	m.TargetModelUri = v
}

func (m *TuningMetadata) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *TuningMetadata) SetStatus(v TuningTask_Status) {
	m.Status = v
}

func (m *TuningMetadata) SetTotalSteps(v int64) {
	m.TotalSteps = v
}

func (m *TuningMetadata) SetCurrentStep(v int64) {
	m.CurrentStep = v
}

type TuningRequest_TuningParams = isTuningRequest_TuningParams

func (m *TuningRequest) SetTuningParams(v TuningRequest_TuningParams) {
	m.TuningParams = v
}

func (m *TuningRequest) SetBaseModelUri(v string) {
	m.BaseModelUri = v
}

func (m *TuningRequest) SetTrainDatasets(v []*TuningRequest_WeightedDataset) {
	m.TrainDatasets = v
}

func (m *TuningRequest) SetValidationDatasets(v []*TuningRequest_WeightedDataset) {
	m.ValidationDatasets = v
}

func (m *TuningRequest) SetTextToTextCompletion(v *TextToTextCompletionTuningParams) {
	m.TuningParams = &TuningRequest_TextToTextCompletion{
		TextToTextCompletion: v,
	}
}

func (m *TuningRequest) SetTextClassificationMultilabel(v *TextClassificationMultilabelParams) {
	m.TuningParams = &TuningRequest_TextClassificationMultilabel{
		TextClassificationMultilabel: v,
	}
}

func (m *TuningRequest) SetTextClassificationMulticlass(v *TextClassificationMulticlassParams) {
	m.TuningParams = &TuningRequest_TextClassificationMulticlass{
		TextClassificationMulticlass: v,
	}
}

func (m *TuningRequest) SetTextEmbeddingPairParams(v *TextEmbeddingPairParams) {
	m.TuningParams = &TuningRequest_TextEmbeddingPairParams{
		TextEmbeddingPairParams: v,
	}
}

func (m *TuningRequest) SetTextEmbeddingTripletParams(v *TextEmbeddingTripletParams) {
	m.TuningParams = &TuningRequest_TextEmbeddingTripletParams{
		TextEmbeddingTripletParams: v,
	}
}

func (m *TuningRequest) SetName(v string) {
	m.Name = v
}

func (m *TuningRequest) SetDescription(v string) {
	m.Description = v
}

func (m *TuningRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *TuningRequest_WeightedDataset) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *TuningRequest_WeightedDataset) SetWeight(v float64) {
	m.Weight = v
}

type TextToTextCompletionTuningParams_TuningType = isTextToTextCompletionTuningParams_TuningType

func (m *TextToTextCompletionTuningParams) SetTuningType(v TextToTextCompletionTuningParams_TuningType) {
	m.TuningType = v
}

func (m *TextToTextCompletionTuningParams) SetSeed(v int64) {
	m.Seed = v
}

func (m *TextToTextCompletionTuningParams) SetLr(v float64) {
	m.Lr = v
}

func (m *TextToTextCompletionTuningParams) SetNSamples(v int64) {
	m.NSamples = v
}

func (m *TextToTextCompletionTuningParams) SetAdditionalArguments(v string) {
	m.AdditionalArguments = v
}

func (m *TextToTextCompletionTuningParams) SetLora(v *TuningTypeLora) {
	m.TuningType = &TextToTextCompletionTuningParams_Lora{
		Lora: v,
	}
}

func (m *TextToTextCompletionTuningParams) SetPromptTune(v *TuningTypePromptTune) {
	m.TuningType = &TextToTextCompletionTuningParams_PromptTune{
		PromptTune: v,
	}
}

func (m *TextToTextCompletionTuningParams) SetScheduler(v *TextToTextCompletionTuningParams_Scheduler) {
	m.Scheduler = v
}

func (m *TextToTextCompletionTuningParams) SetOptimizer(v *TextToTextCompletionTuningParams_Optimizer) {
	m.Optimizer = v
}

type TextToTextCompletionTuningParams_Scheduler_Type = isTextToTextCompletionTuningParams_Scheduler_Type

func (m *TextToTextCompletionTuningParams_Scheduler) SetType(v TextToTextCompletionTuningParams_Scheduler_Type) {
	m.Type = v
}

func (m *TextToTextCompletionTuningParams_Scheduler) SetLinear(v *SchedulerLinear) {
	m.Type = &TextToTextCompletionTuningParams_Scheduler_Linear{
		Linear: v,
	}
}

func (m *TextToTextCompletionTuningParams_Scheduler) SetConstant(v *SchedulerConstant) {
	m.Type = &TextToTextCompletionTuningParams_Scheduler_Constant{
		Constant: v,
	}
}

func (m *TextToTextCompletionTuningParams_Scheduler) SetCosine(v *SchedulerCosine) {
	m.Type = &TextToTextCompletionTuningParams_Scheduler_Cosine{
		Cosine: v,
	}
}

func (m *TextToTextCompletionTuningParams_Scheduler) SetWarmupRatio(v float64) {
	m.WarmupRatio = v
}

type TextToTextCompletionTuningParams_Optimizer_Type = isTextToTextCompletionTuningParams_Optimizer_Type

func (m *TextToTextCompletionTuningParams_Optimizer) SetType(v TextToTextCompletionTuningParams_Optimizer_Type) {
	m.Type = v
}

func (m *TextToTextCompletionTuningParams_Optimizer) SetAdamw(v *OptimizerAdamw) {
	m.Type = &TextToTextCompletionTuningParams_Optimizer_Adamw{
		Adamw: v,
	}
}

type TextClassificationMultilabelParams_TuningType = isTextClassificationMultilabelParams_TuningType

func (m *TextClassificationMultilabelParams) SetTuningType(v TextClassificationMultilabelParams_TuningType) {
	m.TuningType = v
}

func (m *TextClassificationMultilabelParams) SetSeed(v int64) {
	m.Seed = v
}

func (m *TextClassificationMultilabelParams) SetLr(v float64) {
	m.Lr = v
}

func (m *TextClassificationMultilabelParams) SetNSamples(v int64) {
	m.NSamples = v
}

func (m *TextClassificationMultilabelParams) SetAdditionalArguments(v string) {
	m.AdditionalArguments = v
}

func (m *TextClassificationMultilabelParams) SetLora(v *TuningTypeLora) {
	m.TuningType = &TextClassificationMultilabelParams_Lora{
		Lora: v,
	}
}

func (m *TextClassificationMultilabelParams) SetPromptTune(v *TuningTypePromptTune) {
	m.TuningType = &TextClassificationMultilabelParams_PromptTune{
		PromptTune: v,
	}
}

func (m *TextClassificationMultilabelParams) SetScheduler(v *TextClassificationMultilabelParams_Scheduler) {
	m.Scheduler = v
}

func (m *TextClassificationMultilabelParams) SetOptimizer(v *TextClassificationMultilabelParams_Optimizer) {
	m.Optimizer = v
}

type TextClassificationMultilabelParams_Scheduler_Type = isTextClassificationMultilabelParams_Scheduler_Type

func (m *TextClassificationMultilabelParams_Scheduler) SetType(v TextClassificationMultilabelParams_Scheduler_Type) {
	m.Type = v
}

func (m *TextClassificationMultilabelParams_Scheduler) SetLinear(v *SchedulerLinear) {
	m.Type = &TextClassificationMultilabelParams_Scheduler_Linear{
		Linear: v,
	}
}

func (m *TextClassificationMultilabelParams_Scheduler) SetConstant(v *SchedulerConstant) {
	m.Type = &TextClassificationMultilabelParams_Scheduler_Constant{
		Constant: v,
	}
}

func (m *TextClassificationMultilabelParams_Scheduler) SetCosine(v *SchedulerCosine) {
	m.Type = &TextClassificationMultilabelParams_Scheduler_Cosine{
		Cosine: v,
	}
}

func (m *TextClassificationMultilabelParams_Scheduler) SetWarmupRatio(v float64) {
	m.WarmupRatio = v
}

type TextClassificationMultilabelParams_Optimizer_Type = isTextClassificationMultilabelParams_Optimizer_Type

func (m *TextClassificationMultilabelParams_Optimizer) SetType(v TextClassificationMultilabelParams_Optimizer_Type) {
	m.Type = v
}

func (m *TextClassificationMultilabelParams_Optimizer) SetAdamw(v *OptimizerAdamw) {
	m.Type = &TextClassificationMultilabelParams_Optimizer_Adamw{
		Adamw: v,
	}
}

type TextClassificationMulticlassParams_TuningType = isTextClassificationMulticlassParams_TuningType

func (m *TextClassificationMulticlassParams) SetTuningType(v TextClassificationMulticlassParams_TuningType) {
	m.TuningType = v
}

func (m *TextClassificationMulticlassParams) SetSeed(v int64) {
	m.Seed = v
}

func (m *TextClassificationMulticlassParams) SetLr(v float64) {
	m.Lr = v
}

func (m *TextClassificationMulticlassParams) SetNSamples(v int64) {
	m.NSamples = v
}

func (m *TextClassificationMulticlassParams) SetAdditionalArguments(v string) {
	m.AdditionalArguments = v
}

func (m *TextClassificationMulticlassParams) SetLora(v *TuningTypeLora) {
	m.TuningType = &TextClassificationMulticlassParams_Lora{
		Lora: v,
	}
}

func (m *TextClassificationMulticlassParams) SetPromptTune(v *TuningTypePromptTune) {
	m.TuningType = &TextClassificationMulticlassParams_PromptTune{
		PromptTune: v,
	}
}

func (m *TextClassificationMulticlassParams) SetScheduler(v *TextClassificationMulticlassParams_Scheduler) {
	m.Scheduler = v
}

func (m *TextClassificationMulticlassParams) SetOptimizer(v *TextClassificationMulticlassParams_Optimizer) {
	m.Optimizer = v
}

type TextClassificationMulticlassParams_Scheduler_Type = isTextClassificationMulticlassParams_Scheduler_Type

func (m *TextClassificationMulticlassParams_Scheduler) SetType(v TextClassificationMulticlassParams_Scheduler_Type) {
	m.Type = v
}

func (m *TextClassificationMulticlassParams_Scheduler) SetLinear(v *SchedulerLinear) {
	m.Type = &TextClassificationMulticlassParams_Scheduler_Linear{
		Linear: v,
	}
}

func (m *TextClassificationMulticlassParams_Scheduler) SetConstant(v *SchedulerConstant) {
	m.Type = &TextClassificationMulticlassParams_Scheduler_Constant{
		Constant: v,
	}
}

func (m *TextClassificationMulticlassParams_Scheduler) SetCosine(v *SchedulerCosine) {
	m.Type = &TextClassificationMulticlassParams_Scheduler_Cosine{
		Cosine: v,
	}
}

func (m *TextClassificationMulticlassParams_Scheduler) SetWarmupRatio(v float64) {
	m.WarmupRatio = v
}

type TextClassificationMulticlassParams_Optimizer_Type = isTextClassificationMulticlassParams_Optimizer_Type

func (m *TextClassificationMulticlassParams_Optimizer) SetType(v TextClassificationMulticlassParams_Optimizer_Type) {
	m.Type = v
}

func (m *TextClassificationMulticlassParams_Optimizer) SetAdamw(v *OptimizerAdamw) {
	m.Type = &TextClassificationMulticlassParams_Optimizer_Adamw{
		Adamw: v,
	}
}

type TextEmbeddingPairParams_TuningType = isTextEmbeddingPairParams_TuningType

func (m *TextEmbeddingPairParams) SetTuningType(v TextEmbeddingPairParams_TuningType) {
	m.TuningType = v
}

func (m *TextEmbeddingPairParams) SetSeed(v int64) {
	m.Seed = v
}

func (m *TextEmbeddingPairParams) SetLr(v float64) {
	m.Lr = v
}

func (m *TextEmbeddingPairParams) SetNSamples(v int64) {
	m.NSamples = v
}

func (m *TextEmbeddingPairParams) SetAdditionalArguments(v string) {
	m.AdditionalArguments = v
}

func (m *TextEmbeddingPairParams) SetEmbeddingDims(v []int64) {
	m.EmbeddingDims = v
}

func (m *TextEmbeddingPairParams) SetLora(v *TuningTypeLora) {
	m.TuningType = &TextEmbeddingPairParams_Lora{
		Lora: v,
	}
}

func (m *TextEmbeddingPairParams) SetPromptTune(v *TuningTypePromptTune) {
	m.TuningType = &TextEmbeddingPairParams_PromptTune{
		PromptTune: v,
	}
}

func (m *TextEmbeddingPairParams) SetScheduler(v *TextEmbeddingPairParams_Scheduler) {
	m.Scheduler = v
}

func (m *TextEmbeddingPairParams) SetOptimizer(v *TextEmbeddingPairParams_Optimizer) {
	m.Optimizer = v
}

type TextEmbeddingPairParams_Scheduler_Type = isTextEmbeddingPairParams_Scheduler_Type

func (m *TextEmbeddingPairParams_Scheduler) SetType(v TextEmbeddingPairParams_Scheduler_Type) {
	m.Type = v
}

type TextEmbeddingPairParams_Scheduler_OptionalWarmupRatio = isTextEmbeddingPairParams_Scheduler_OptionalWarmupRatio

func (m *TextEmbeddingPairParams_Scheduler) SetOptionalWarmupRatio(v TextEmbeddingPairParams_Scheduler_OptionalWarmupRatio) {
	m.OptionalWarmupRatio = v
}

func (m *TextEmbeddingPairParams_Scheduler) SetLinear(v *SchedulerLinear) {
	m.Type = &TextEmbeddingPairParams_Scheduler_Linear{
		Linear: v,
	}
}

func (m *TextEmbeddingPairParams_Scheduler) SetConstant(v *SchedulerConstant) {
	m.Type = &TextEmbeddingPairParams_Scheduler_Constant{
		Constant: v,
	}
}

func (m *TextEmbeddingPairParams_Scheduler) SetCosine(v *SchedulerCosine) {
	m.Type = &TextEmbeddingPairParams_Scheduler_Cosine{
		Cosine: v,
	}
}

func (m *TextEmbeddingPairParams_Scheduler) SetWarmupRatio(v float64) {
	m.OptionalWarmupRatio = &TextEmbeddingPairParams_Scheduler_WarmupRatio{
		WarmupRatio: v,
	}
}

type TextEmbeddingPairParams_Optimizer_Type = isTextEmbeddingPairParams_Optimizer_Type

func (m *TextEmbeddingPairParams_Optimizer) SetType(v TextEmbeddingPairParams_Optimizer_Type) {
	m.Type = v
}

func (m *TextEmbeddingPairParams_Optimizer) SetAdamw(v *OptimizerAdamw) {
	m.Type = &TextEmbeddingPairParams_Optimizer_Adamw{
		Adamw: v,
	}
}

type TextEmbeddingTripletParams_TuningType = isTextEmbeddingTripletParams_TuningType

func (m *TextEmbeddingTripletParams) SetTuningType(v TextEmbeddingTripletParams_TuningType) {
	m.TuningType = v
}

func (m *TextEmbeddingTripletParams) SetSeed(v int64) {
	m.Seed = v
}

func (m *TextEmbeddingTripletParams) SetLr(v float64) {
	m.Lr = v
}

func (m *TextEmbeddingTripletParams) SetNSamples(v int64) {
	m.NSamples = v
}

func (m *TextEmbeddingTripletParams) SetAdditionalArguments(v string) {
	m.AdditionalArguments = v
}

func (m *TextEmbeddingTripletParams) SetEmbeddingDims(v []int64) {
	m.EmbeddingDims = v
}

func (m *TextEmbeddingTripletParams) SetLora(v *TuningTypeLora) {
	m.TuningType = &TextEmbeddingTripletParams_Lora{
		Lora: v,
	}
}

func (m *TextEmbeddingTripletParams) SetPromptTune(v *TuningTypePromptTune) {
	m.TuningType = &TextEmbeddingTripletParams_PromptTune{
		PromptTune: v,
	}
}

func (m *TextEmbeddingTripletParams) SetScheduler(v *TextEmbeddingTripletParams_Scheduler) {
	m.Scheduler = v
}

func (m *TextEmbeddingTripletParams) SetOptimizer(v *TextEmbeddingTripletParams_Optimizer) {
	m.Optimizer = v
}

type TextEmbeddingTripletParams_Scheduler_Type = isTextEmbeddingTripletParams_Scheduler_Type

func (m *TextEmbeddingTripletParams_Scheduler) SetType(v TextEmbeddingTripletParams_Scheduler_Type) {
	m.Type = v
}

type TextEmbeddingTripletParams_Scheduler_OptionalWarmupRatio = isTextEmbeddingTripletParams_Scheduler_OptionalWarmupRatio

func (m *TextEmbeddingTripletParams_Scheduler) SetOptionalWarmupRatio(v TextEmbeddingTripletParams_Scheduler_OptionalWarmupRatio) {
	m.OptionalWarmupRatio = v
}

func (m *TextEmbeddingTripletParams_Scheduler) SetLinear(v *SchedulerLinear) {
	m.Type = &TextEmbeddingTripletParams_Scheduler_Linear{
		Linear: v,
	}
}

func (m *TextEmbeddingTripletParams_Scheduler) SetConstant(v *SchedulerConstant) {
	m.Type = &TextEmbeddingTripletParams_Scheduler_Constant{
		Constant: v,
	}
}

func (m *TextEmbeddingTripletParams_Scheduler) SetCosine(v *SchedulerCosine) {
	m.Type = &TextEmbeddingTripletParams_Scheduler_Cosine{
		Cosine: v,
	}
}

func (m *TextEmbeddingTripletParams_Scheduler) SetWarmupRatio(v float64) {
	m.OptionalWarmupRatio = &TextEmbeddingTripletParams_Scheduler_WarmupRatio{
		WarmupRatio: v,
	}
}

type TextEmbeddingTripletParams_Optimizer_Type = isTextEmbeddingTripletParams_Optimizer_Type

func (m *TextEmbeddingTripletParams_Optimizer) SetType(v TextEmbeddingTripletParams_Optimizer_Type) {
	m.Type = v
}

func (m *TextEmbeddingTripletParams_Optimizer) SetAdamw(v *OptimizerAdamw) {
	m.Type = &TextEmbeddingTripletParams_Optimizer_Adamw{
		Adamw: v,
	}
}

func (m *GetMetricsUrlRequest) SetTaskId(v string) {
	m.TaskId = v
}

func (m *GetMetricsUrlResponse) SetLoadUrl(v string) {
	m.LoadUrl = v
}

func (m *GetOptionsRequest) SetTaskId(v string) {
	m.TaskId = v
}

type GetOptionsResponse_TuningParams = isGetOptionsResponse_TuningParams

func (m *GetOptionsResponse) SetTuningParams(v GetOptionsResponse_TuningParams) {
	m.TuningParams = v
}

func (m *GetOptionsResponse) SetTaskId(v string) {
	m.TaskId = v
}

func (m *GetOptionsResponse) SetBaseModelUri(v string) {
	m.BaseModelUri = v
}

func (m *GetOptionsResponse) SetTrainDatasets(v []*TuningRequest_WeightedDataset) {
	m.TrainDatasets = v
}

func (m *GetOptionsResponse) SetValidationDatasets(v []*TuningRequest_WeightedDataset) {
	m.ValidationDatasets = v
}

func (m *GetOptionsResponse) SetTextToTextCompletion(v *TextToTextCompletionTuningParams) {
	m.TuningParams = &GetOptionsResponse_TextToTextCompletion{
		TextToTextCompletion: v,
	}
}

func (m *GetOptionsResponse) SetTextClassificationMultilabel(v *TextClassificationMultilabelParams) {
	m.TuningParams = &GetOptionsResponse_TextClassificationMultilabel{
		TextClassificationMultilabel: v,
	}
}

func (m *GetOptionsResponse) SetTextClassificationMulticlass(v *TextClassificationMulticlassParams) {
	m.TuningParams = &GetOptionsResponse_TextClassificationMulticlass{
		TextClassificationMulticlass: v,
	}
}

func (m *GetOptionsResponse) SetTextEmbeddingPairParams(v *TextEmbeddingPairParams) {
	m.TuningParams = &GetOptionsResponse_TextEmbeddingPairParams{
		TextEmbeddingPairParams: v,
	}
}

func (m *GetOptionsResponse) SetTextEmbeddingTripletParams(v *TextEmbeddingTripletParams) {
	m.TuningParams = &GetOptionsResponse_TextEmbeddingTripletParams{
		TextEmbeddingTripletParams: v,
	}
}

func (m *ListErrorsRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *ListErrorsResponse) SetTuningError(v []*TuningError) {
	m.TuningError = v
}

type CreateTuningDraftRequest_TuningParams = isCreateTuningDraftRequest_TuningParams

func (m *CreateTuningDraftRequest) SetTuningParams(v CreateTuningDraftRequest_TuningParams) {
	m.TuningParams = v
}

func (m *CreateTuningDraftRequest) SetBaseModelUri(v string) {
	m.BaseModelUri = v
}

func (m *CreateTuningDraftRequest) SetTrainDatasets(v []*TuningRequest_WeightedDataset) {
	m.TrainDatasets = v
}

func (m *CreateTuningDraftRequest) SetValidationDatasets(v []*TuningRequest_WeightedDataset) {
	m.ValidationDatasets = v
}

func (m *CreateTuningDraftRequest) SetTextToTextCompletion(v *TextToTextCompletionTuningParams) {
	m.TuningParams = &CreateTuningDraftRequest_TextToTextCompletion{
		TextToTextCompletion: v,
	}
}

func (m *CreateTuningDraftRequest) SetTextClassificationMultilabel(v *TextClassificationMultilabelParams) {
	m.TuningParams = &CreateTuningDraftRequest_TextClassificationMultilabel{
		TextClassificationMultilabel: v,
	}
}

func (m *CreateTuningDraftRequest) SetTextClassificationMulticlass(v *TextClassificationMulticlassParams) {
	m.TuningParams = &CreateTuningDraftRequest_TextClassificationMulticlass{
		TextClassificationMulticlass: v,
	}
}

func (m *CreateTuningDraftRequest) SetTextEmbeddingPairParams(v *TextEmbeddingPairParams) {
	m.TuningParams = &CreateTuningDraftRequest_TextEmbeddingPairParams{
		TextEmbeddingPairParams: v,
	}
}

func (m *CreateTuningDraftRequest) SetTextEmbeddingTripletParams(v *TextEmbeddingTripletParams) {
	m.TuningParams = &CreateTuningDraftRequest_TextEmbeddingTripletParams{
		TextEmbeddingTripletParams: v,
	}
}

func (m *CreateTuningDraftRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateTuningDraftRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateTuningDraftRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateTuningDraftResponse) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

type UpdateTuningDraftRequest_TuningParams = isUpdateTuningDraftRequest_TuningParams

func (m *UpdateTuningDraftRequest) SetTuningParams(v UpdateTuningDraftRequest_TuningParams) {
	m.TuningParams = v
}

func (m *UpdateTuningDraftRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *UpdateTuningDraftRequest) SetBaseModelUri(v string) {
	m.BaseModelUri = v
}

func (m *UpdateTuningDraftRequest) SetTrainDatasets(v []*TuningRequest_WeightedDataset) {
	m.TrainDatasets = v
}

func (m *UpdateTuningDraftRequest) SetValidationDatasets(v []*TuningRequest_WeightedDataset) {
	m.ValidationDatasets = v
}

func (m *UpdateTuningDraftRequest) SetTextToTextCompletion(v *TextToTextCompletionTuningParams) {
	m.TuningParams = &UpdateTuningDraftRequest_TextToTextCompletion{
		TextToTextCompletion: v,
	}
}

func (m *UpdateTuningDraftRequest) SetTextClassificationMultilabel(v *TextClassificationMultilabelParams) {
	m.TuningParams = &UpdateTuningDraftRequest_TextClassificationMultilabel{
		TextClassificationMultilabel: v,
	}
}

func (m *UpdateTuningDraftRequest) SetTextClassificationMulticlass(v *TextClassificationMulticlassParams) {
	m.TuningParams = &UpdateTuningDraftRequest_TextClassificationMulticlass{
		TextClassificationMulticlass: v,
	}
}

func (m *UpdateTuningDraftRequest) SetTextEmbeddingPairParams(v *TextEmbeddingPairParams) {
	m.TuningParams = &UpdateTuningDraftRequest_TextEmbeddingPairParams{
		TextEmbeddingPairParams: v,
	}
}

func (m *UpdateTuningDraftRequest) SetTextEmbeddingTripletParams(v *TextEmbeddingTripletParams) {
	m.TuningParams = &UpdateTuningDraftRequest_TextEmbeddingTripletParams{
		TextEmbeddingTripletParams: v,
	}
}

func (m *UpdateTuningDraftRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateTuningDraftRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateTuningDraftRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateTuningDraftResponse) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *DeleteTuningDraftRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *DeleteTuningDraftResponse) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}

func (m *TuneDraftRequest) SetTuningTaskId(v string) {
	m.TuningTaskId = v
}
