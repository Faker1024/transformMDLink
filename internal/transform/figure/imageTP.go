package figure

import "transformMDLink/internal/transform/utils"

type imageTP struct {
	token string
}

func (i *imageTP) upload(path *string) {
	_, err := utils.ReadImage(*path)
	if err != nil {
		return
	}

}
