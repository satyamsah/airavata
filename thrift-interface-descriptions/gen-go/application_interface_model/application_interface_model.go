// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package application_interface_model

import (
	"bytes"
	"reflect"
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"application_io_models"
	"airavata_commons"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

var _ = application_io_models.GoUnusedProtection__
var _ = airavata_commons.GoUnusedProtection__
// Application Interface Description
// 
// applicationModules:
//   Associate all application modules with versions which interface is applicable to.
// 
// applicationInputs:
//   Inputs to be passed to the application
// 
// applicationOutputs:
//   Outputs generated from the application
// 
// 
// Attributes:
//  - ApplicationInterfaceId
//  - ApplicationName
//  - ApplicationDescription
//  - ApplicationModules
//  - ApplicationInputs
//  - ApplicationOutputs
//  - ArchiveWorkingDirectory
//  - HasOptionalFileInputs
type ApplicationInterfaceDescription struct {
  ApplicationInterfaceId string `thrift:"applicationInterfaceId,1,required" db:"applicationInterfaceId" json:"applicationInterfaceId"`
  ApplicationName string `thrift:"applicationName,2,required" db:"applicationName" json:"applicationName"`
  ApplicationDescription *string `thrift:"applicationDescription,3" db:"applicationDescription" json:"applicationDescription,omitempty"`
  ApplicationModules []string `thrift:"applicationModules,4" db:"applicationModules" json:"applicationModules,omitempty"`
  ApplicationInputs []*application_io_models.InputDataObjectType `thrift:"applicationInputs,5" db:"applicationInputs" json:"applicationInputs,omitempty"`
  ApplicationOutputs []*application_io_models.OutputDataObjectType `thrift:"applicationOutputs,6" db:"applicationOutputs" json:"applicationOutputs,omitempty"`
  ArchiveWorkingDirectory bool `thrift:"archiveWorkingDirectory,7" db:"archiveWorkingDirectory" json:"archiveWorkingDirectory,omitempty"`
  HasOptionalFileInputs *bool `thrift:"hasOptionalFileInputs,8" db:"hasOptionalFileInputs" json:"hasOptionalFileInputs,omitempty"`
}

func NewApplicationInterfaceDescription() *ApplicationInterfaceDescription {
  return &ApplicationInterfaceDescription{
ApplicationInterfaceId: "DO_NOT_SET_AT_CLIENTS",
}
}


func (p *ApplicationInterfaceDescription) GetApplicationInterfaceId() string {
  return p.ApplicationInterfaceId
}

func (p *ApplicationInterfaceDescription) GetApplicationName() string {
  return p.ApplicationName
}
var ApplicationInterfaceDescription_ApplicationDescription_DEFAULT string
func (p *ApplicationInterfaceDescription) GetApplicationDescription() string {
  if !p.IsSetApplicationDescription() {
    return ApplicationInterfaceDescription_ApplicationDescription_DEFAULT
  }
return *p.ApplicationDescription
}
var ApplicationInterfaceDescription_ApplicationModules_DEFAULT []string

func (p *ApplicationInterfaceDescription) GetApplicationModules() []string {
  return p.ApplicationModules
}
var ApplicationInterfaceDescription_ApplicationInputs_DEFAULT []*application_io_models.InputDataObjectType

func (p *ApplicationInterfaceDescription) GetApplicationInputs() []*application_io_models.InputDataObjectType {
  return p.ApplicationInputs
}
var ApplicationInterfaceDescription_ApplicationOutputs_DEFAULT []*application_io_models.OutputDataObjectType

func (p *ApplicationInterfaceDescription) GetApplicationOutputs() []*application_io_models.OutputDataObjectType {
  return p.ApplicationOutputs
}
var ApplicationInterfaceDescription_ArchiveWorkingDirectory_DEFAULT bool = false

func (p *ApplicationInterfaceDescription) GetArchiveWorkingDirectory() bool {
  return p.ArchiveWorkingDirectory
}
var ApplicationInterfaceDescription_HasOptionalFileInputs_DEFAULT bool
func (p *ApplicationInterfaceDescription) GetHasOptionalFileInputs() bool {
  if !p.IsSetHasOptionalFileInputs() {
    return ApplicationInterfaceDescription_HasOptionalFileInputs_DEFAULT
  }
return *p.HasOptionalFileInputs
}
func (p *ApplicationInterfaceDescription) IsSetApplicationDescription() bool {
  return p.ApplicationDescription != nil
}

func (p *ApplicationInterfaceDescription) IsSetApplicationModules() bool {
  return p.ApplicationModules != nil
}

func (p *ApplicationInterfaceDescription) IsSetApplicationInputs() bool {
  return p.ApplicationInputs != nil
}

func (p *ApplicationInterfaceDescription) IsSetApplicationOutputs() bool {
  return p.ApplicationOutputs != nil
}

func (p *ApplicationInterfaceDescription) IsSetArchiveWorkingDirectory() bool {
  return p.ArchiveWorkingDirectory != ApplicationInterfaceDescription_ArchiveWorkingDirectory_DEFAULT
}

func (p *ApplicationInterfaceDescription) IsSetHasOptionalFileInputs() bool {
  return p.HasOptionalFileInputs != nil
}

func (p *ApplicationInterfaceDescription) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }

  var issetApplicationInterfaceId bool = false;
  var issetApplicationName bool = false;

  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetApplicationInterfaceId = true
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
      issetApplicationName = true
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 7:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField7(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 8:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField8(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  if !issetApplicationInterfaceId{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ApplicationInterfaceId is not set"));
  }
  if !issetApplicationName{
    return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ApplicationName is not set"));
  }
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.ApplicationInterfaceId = v
}
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.ApplicationName = v
}
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.ApplicationDescription = &v
}
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField4(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]string, 0, size)
  p.ApplicationModules =  tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return thrift.PrependError("error reading field 0: ", err)
} else {
    _elem0 = v
}
    p.ApplicationModules = append(p.ApplicationModules, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField5(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*application_io_models.InputDataObjectType, 0, size)
  p.ApplicationInputs =  tSlice
  for i := 0; i < size; i ++ {
    _elem1 := &application_io_models.InputDataObjectType{}
    if err := _elem1.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem1), err)
    }
    p.ApplicationInputs = append(p.ApplicationInputs, _elem1)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField6(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*application_io_models.OutputDataObjectType, 0, size)
  p.ApplicationOutputs =  tSlice
  for i := 0; i < size; i ++ {
    _elem2 := &application_io_models.OutputDataObjectType{}
    if err := _elem2.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem2), err)
    }
    p.ApplicationOutputs = append(p.ApplicationOutputs, _elem2)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField7(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 7: ", err)
} else {
  p.ArchiveWorkingDirectory = v
}
  return nil
}

func (p *ApplicationInterfaceDescription)  ReadField8(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 8: ", err)
} else {
  p.HasOptionalFileInputs = &v
}
  return nil
}

func (p *ApplicationInterfaceDescription) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("ApplicationInterfaceDescription"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField7(oprot); err != nil { return err }
    if err := p.writeField8(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *ApplicationInterfaceDescription) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("applicationInterfaceId", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:applicationInterfaceId: ", p), err) }
  if err := oprot.WriteString(string(p.ApplicationInterfaceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.applicationInterfaceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:applicationInterfaceId: ", p), err) }
  return err
}

func (p *ApplicationInterfaceDescription) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("applicationName", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:applicationName: ", p), err) }
  if err := oprot.WriteString(string(p.ApplicationName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.applicationName (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:applicationName: ", p), err) }
  return err
}

func (p *ApplicationInterfaceDescription) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetApplicationDescription() {
    if err := oprot.WriteFieldBegin("applicationDescription", thrift.STRING, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:applicationDescription: ", p), err) }
    if err := oprot.WriteString(string(*p.ApplicationDescription)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.applicationDescription (3) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:applicationDescription: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetApplicationModules() {
    if err := oprot.WriteFieldBegin("applicationModules", thrift.LIST, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:applicationModules: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRING, len(p.ApplicationModules)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.ApplicationModules {
      if err := oprot.WriteString(string(v)); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:applicationModules: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetApplicationInputs() {
    if err := oprot.WriteFieldBegin("applicationInputs", thrift.LIST, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:applicationInputs: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.ApplicationInputs)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.ApplicationInputs {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:applicationInputs: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) writeField6(oprot thrift.TProtocol) (err error) {
  if p.IsSetApplicationOutputs() {
    if err := oprot.WriteFieldBegin("applicationOutputs", thrift.LIST, 6); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:applicationOutputs: ", p), err) }
    if err := oprot.WriteListBegin(thrift.STRUCT, len(p.ApplicationOutputs)); err != nil {
      return thrift.PrependError("error writing list begin: ", err)
    }
    for _, v := range p.ApplicationOutputs {
      if err := v.Write(oprot); err != nil {
        return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
      }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return thrift.PrependError("error writing list end: ", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 6:applicationOutputs: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) writeField7(oprot thrift.TProtocol) (err error) {
  if p.IsSetArchiveWorkingDirectory() {
    if err := oprot.WriteFieldBegin("archiveWorkingDirectory", thrift.BOOL, 7); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:archiveWorkingDirectory: ", p), err) }
    if err := oprot.WriteBool(bool(p.ArchiveWorkingDirectory)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.archiveWorkingDirectory (7) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 7:archiveWorkingDirectory: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) writeField8(oprot thrift.TProtocol) (err error) {
  if p.IsSetHasOptionalFileInputs() {
    if err := oprot.WriteFieldBegin("hasOptionalFileInputs", thrift.BOOL, 8); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:hasOptionalFileInputs: ", p), err) }
    if err := oprot.WriteBool(bool(*p.HasOptionalFileInputs)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.hasOptionalFileInputs (8) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 8:hasOptionalFileInputs: ", p), err) }
  }
  return err
}

func (p *ApplicationInterfaceDescription) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ApplicationInterfaceDescription(%+v)", *p)
}
