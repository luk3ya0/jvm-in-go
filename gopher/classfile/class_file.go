package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse construct ClassFile and ClassReader
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)

	return
}

// read is a workflow of loading items of ClassFile from ClassReader
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader) // items of u1, u2, u4 type
	self.accessFlags = reader.readUint16()       // 16 bit mask for "class or interface", "private or public", etc.
	self.thisClass = reader.readUint16()         // index in constant pool
	self.superClass = reader.readUint16()        // index in constant pool
	self.interfaces = reader.readUint16s()       // indexes in constant pool
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

// readAndCheckMagic grasp the first four byte of data to check
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic")
	}
}

// readAndCheckVersion grasp the first four byte of data to load major and minor version
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
}

// MinorVersion a exported getter for minor version
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

// MajorVersion a exported getter for major version
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// ConstantPool a exported getter for constant pool
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

// AccessFlags a exported getter for access flag
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

// Fields a exported getter for fields table
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

// Methods a exported getter for methods table
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

// ClassName a exported getter for class name from constant pool
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName a exported getter for super class name from constant pool
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}

	return "" // only java.lang.Object has not super class
}

// InterfaceNames a exported getter for interfaces table
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}
