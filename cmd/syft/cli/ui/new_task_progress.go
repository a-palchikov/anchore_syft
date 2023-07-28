package ui

import "github.com/anchore/bubbly/bubbles/taskprogress"

func (m *Handler) newTaskProgress(title taskprogress.Title, opts ...taskprogress.Option) *taskprogress.Model {
	opts = append(opts, taskprogress.WithName(title.Default))
	tsk := taskprogress.New(m.Running, opts...)

	tsk.HideProgressOnSuccess = true
	tsk.HideStageOnSuccess = true
	tsk.WindowSize = m.WindowSize
	tsk.TitleWidth = m.Config.TitleWidth
	tsk.TitleOptions = title

	if m.Config.AdjustDefaultTask != nil {
		tsk = m.Config.AdjustDefaultTask(tsk)
	}

	return tsk
}
