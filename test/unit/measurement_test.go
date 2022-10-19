package unit

import (
	"deepfit/internal/user/measurement"
	"deepfit/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

func TestMeasurement(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "measurement")
}

var _ = Describe("Testing Measurement Package", func() {
	Context("Testing SetId", func() {
		It("should return set id", func() {
			measurementObj := measurement.Measurement{}
			id := primitive.NewObjectID()
			Expect(measurementObj.SetId(id).Id).To(Equal(id))
		})
	})

	Context("Testing SetRandomId", func() {
		It("should return set random id", func() {
			measurementObj := measurement.Measurement{}
			Expect(measurementObj.SetRandomId().Id).ToNot(BeNil())
		})
	})

	Context("Testing SetArm", func() {
		It("should return set arm", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetArm(10)
			Expect(measurementObj.Arm).To(Equal(10))
		})
	})

	Context("Testing SetLeg", func() {
		It("should return set leg", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetLeg(10)
			Expect(measurementObj.Leg).To(Equal(10))
		})
	})

	Context("Testing SetWaist", func() {
		It("should return set waist", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetWaist(10)
			Expect(measurementObj.Waist).To(Equal(10))
		})
	})

	Context("Testing SetWeight", func() {
		It("should return set weight", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetWeight(10)
			Expect(measurementObj.Weight).To(Equal(10))
		})
	})

	Context("Testing SetHeight", func() {
		It("should return set height", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetHeight(10)
			Expect(measurementObj.Height).To(Equal(10))
		})
	})

	Context("Testing SetChest", func() {
		It("should return set chest", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetChest(10)
			Expect(measurementObj.Chest).To(Equal(10))
		})
	})

	Context("Testing SetShoulder", func() {
		It("should return set shoulder", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetShoulder(10)
			Expect(measurementObj.Shoulder).To(Equal(10))
		})
	})

	Context("Testing SetHip", func() {
		It("should return set hip", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetHip(10)
			Expect(measurementObj.Hip).To(Equal(10))
		})
	})

	Context("Testing SetIsPublic", func() {
		It("should return set is public", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetIsPublic(true)
			Expect(measurementObj.IsPublic).To(Equal(true))
		})
	})

	Context("Testing SetDate", func() {
		It("should return set date", func() {
			measurementObj := measurement.Measurement{}
			measurementObj.SetDate()
			Expect(measurementObj.Date.UpdatedAt).ToNot(BeNil())
			Expect(measurementObj.Date.UpdatedAt).To(Equal(time.Time{}))
			Expect(measurementObj.Date.DeletedAt).To(Equal(time.Time{}))
		})
	})

	Context("Testing SetImages", func() {
		It("should return set images", func() {
			measurementObj := measurement.Measurement{}
			images := []pkg.Image{pkg.NewImage("image1")}
			measurementObj.SetImages(images)
			Expect(measurementObj.Images).To(Equal(images))
		})
	})
})
